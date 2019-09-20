package service

import (
	"backend/cache"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Article struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required,max=32"`
	Content     string `json:"content" db:"content" binding:"required"`
	Html        string `json:"html" db:"html" binding:"required"`
	CategoryID  int    `json:"category_id" db:"category_id" binding:"required"`
	TagID       []int  `json:"tag_id" binding:"required"`
	CreatedTime string `json:"created_time" db:"created_time"`
	UpdatedTime string `json:"updated_time" db:"updated_time"`
	Status      string `json:"status" db:"status" binding:"required"`
}

type Articles struct {
	Items []Article `json:"items"`
	Total int       `json:"total"`
}

type ArticleDetail struct {
	A    Article  `json:"article"`
	C    Category `json:"category"`
	Tags []Tag    `json:"tags"`
}

type Options struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Search bool   `json:"search"`
	Admin  bool   `json:"admin"`
	C      string `json:"c"` // category
	T      string `json:"t"` // tag
}

var defaultOptions = Options{
	Limit:  10,
	Page:   1,
	C:      "",
	T:      "",
	Search: false, // 搜索文章结果不进行缓存
	Admin:  false, // 是否是admin页面请求，如果不是，文章就不包括草稿文章
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	// 初始化默认值
	opt := defaultOptions

	for _, o := range opts {
		o(&opt) // 依次调用opts函数列表中的函数，为服务选项（opt变量）赋值
	}

	return opt
}

func SetLimitPageAdmin(limit, page, admin string) Option {
	return func(o *Options) {
		if limit != "" && page != "" {
			p, _ := strconv.Atoi(page)
			l, _ := strconv.Atoi(limit)
			o.Limit = l
			o.Page = p
		}
		if admin != "" {
			o.Admin = true
		}
	}
}

func SetCategory(c string) Option {
	return func(o *Options) {
		o.C = c
	}
}

func SetTag(t string) Option {
	return func(o *Options) {
		o.T = t
	}
}

func SetSearch(search bool) Option {
	return func(o *Options) {
		o.Search = search
		o.Page = defaultOptions.Page // 如果不是在第一页执行的搜索，比如：page=3，有可能会搜不到数据，必须从第一页开始搜索
	}
}

var db = models.DB

func (a *Article) Create() (Article, error) {
	createTime := time.Now().Format(utils.AppInfo.TimeFormat)

	r, e := db.Exec("insert into blog_article (title,content,html,category_id,created_time,status) values (?,?,?,?,?,?)", a.Title, a.Content, a.Html,a.CategoryID, createTime, a.Status)
	if e != nil {
		return Article{}, e
	}
	articleID, _ := r.LastInsertId()
	if len(a.TagID) > 0 {
		for _, tagID := range a.TagID {
			_, e := db.Exec("insert into blog_tag_article (tag_id, article_id) values (?, ?)", tagID, articleID)
			if e != nil {
				return Article{}, e
			}
		}
	}
	return Article{int(articleID), a.Title, a.Content, a.Html,a.CategoryID, a.TagID, createTime, a.UpdatedTime, a.Status}, nil
}

func (a *Article) Delete() error {
	if _, e := db.Exec("delete from blog_tag_article where article_id=?", a.ID); e != nil {
		return e
	}
	if _, e := db.Exec("delete from blog_article where id=?", a.ID); e != nil {
		return e
	}
	return nil
}

func (a *Article) Edit() error {
	updateTime := time.Now().Format(utils.AppInfo.TimeFormat)
	if _, e := db.Exec("update blog_article set title=?,content=?,html=?,category_id=?,updated_time=?,`status`=? where id=?", a.Title, a.Content, a.Html, a.CategoryID, updateTime, a.Status, a.ID); e != nil {
		return e
	}

	if _, e := db.Exec("delete from blog_tag_article where article_id=?", a.ID); e != nil {
		return e
	}

	if len(a.TagID) > 0 {
		for _, tagID := range a.TagID {
			_, e := db.Exec("insert into blog_tag_article (tag_id, article_id) values (?, ?)", tagID, a.ID)
			if e != nil {
				return e
			}
		}
	}

	return nil
}

func (a Article) GetOne() (ArticleDetail, error) {
	var one Article
	if e := db.Get(&one, "select * from blog_article where id=?", a.ID); e != nil {
		return ArticleDetail{}, e
	}
	category, _ := GetCategoryByID(one.CategoryID)
	tags, _ := GetTagsByArticleID(a.ID)
	return ArticleDetail{one, category, tags}, nil
}

func (a Article) GetAll(opts ...Option) (data Articles, err error) {
	baseSql := "select %s from blog_article a"
	data, err = genArticles(baseSql, opts...)
	return
}

func GetArticlesByCategory(name string, opts ...Option) (data Articles, err error) {
	baseSql := "SELECT %s FROM blog_article a INNER JOIN blog_category c ON a.category_id=c.id AND c.category_name=" + "'" + name + "'" + ""
	data, err = genArticles(baseSql, opts...)
	return
}

func GetArticlesByTag(name string, opts ...Option) (data Articles, err error) {
	baseSql := "SELECT %s FROM blog_article a  INNER JOIN blog_tag_article ta ON a.id=ta.article_id INNER JOIN blog_tag t ON ta.tag_id=t.id AND t.tag_name=" + "'" + name + "'" + ""
	data, err = genArticles(baseSql, opts...)
	return
}

func genArticles(baseSql string, opts ...Option) (data Articles, err error) {
	options := newOptions(opts...)
	key := cache.ArticleKey(options.Limit, options.Page, options.Admin, options.C, options.T)
	if !options.Search {
		cacheData, e := getCache(key)
		if e != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 读取缓存失败, %v", time.Now().Format(utils.AppInfo.TimeFormat), e))
		}
		if cacheData.Total != 0 {
			return cacheData, nil
		}
	}

	articles := make([]Article, 0)

	var f string
	if !options.Admin {
		f = " WHERE a.`status`='published'"
	}
	offset := (options.Page - 1) * options.Limit
	selectSql := fmt.Sprintf(baseSql, "a.id, a.title, a.created_time, a.updated_time, a.`status`") + f + fmt.Sprintf(" ORDER BY a.id DESC limit %d offset %d", options.Limit, offset)

	if err = db.Select(&articles, selectSql); err != nil {
		return
	}

	var total int
	if err = db.Get(&total, fmt.Sprintf(baseSql, "count(1)")+f); err != nil {
		return
	}

	data.Total = total
	data.Items = articles

	if !options.Search {
		if e := setCache(key, data); e != nil {
			utils.WriteErrorLog(fmt.Sprintf("[ %s ] 写入缓存失败, %v", time.Now().Format(utils.AppInfo.TimeFormat), e))
		}
	}

	return
}

func GetTagsByArticleID(articleID int) ([]Tag, error) {
	var t []Tag
	if e := db.Select(&t, "SELECT t.* FROM blog_tag t RIGHT JOIN blog_tag_article ta ON t.id=ta.tag_id WHERE ta.article_id=?", articleID); e != nil {
		return nil, e
	}
	return t, nil
}

func GetCategoryByID(id int) (Category, error) {
	var c Category
	if e := db.Get(&c, "SELECT * FROM blog_category WHERE id=?", id); e != nil {
		return Category{}, e
	}
	return c, nil
}

func SearchArticle(key, status string, opts ...Option) (data Articles, err error) {
	var baseSql string
	if status == "" {
		baseSql = `SELECT %s FROM blog_article a WHERE a.title LIKE '%%` + key + `%%'`
	} else {
		baseSql = `SELECT %s FROM blog_article a WHERE a.title LIKE '%%` + key + `%%' AND a.status='` + status + `'`
	}

	data, err = genArticles(baseSql, opts...)
	return
}

func getCache(key string) (Articles, error) {
	a := Articles{}
	data, err := cache.GetKey(key)
	if err != nil || data == nil {
		return a, err
	}

	v, ok := data.([]uint8)
	if ok {
		if e := json.Unmarshal([]byte(v[:]), &a); e != nil {
			return a, e
		}
		return a, nil
	} else {
		return a, errors.New("返回数据类型有误，json无法解析")
	}
}

func setCache(key string, value Articles) error {
	marshal, _ := json.Marshal(value)
	e := cache.SetKey(key, marshal)
	return e
}
