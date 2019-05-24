package models

import (
	"backend/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type obj interface{}

type Tag struct {
	ID      int    `json:"id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name" binding:"required,max=16"`
}

type Category struct {
	ID           int    `json:"id" db:"id"`
	CategoryName string `json:"category_name" db:"category_name" binding:"required,max=16"`
}

type Article struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required,max=32"`
	Content     string `json:"content" db:"content" binding:"required"`
	CategoryID  int    `json:"category_id" db:"category_id" binding:"required"`
	TagID       []int  `json:"tag_id" binding:"required"`
	CreatedTime string `json:"created_time" db:"created_time"`
	UpdatedTime string `json:"updated_time" db:"updated_time"`
}

type Base interface {
	Add() (obj, error)
	Delete() error
	Edit() error
	GetOne() (obj, error)
	IfExistByName() bool
}

func HandleRollback(tx *sqlx.Tx) {
	if e := tx.Rollback(); e != nil {
		panic(e)
	}
}

func HandleCommit(tx *sqlx.Tx) {
	if e := tx.Commit(); e != nil {
		HandleRollback(tx)
		panic(e)
	}
}

// 所有tag的api方法
func (t *Tag) IfExistByName() bool {
	tag := new(Tag)
	if e := db.Get(tag, "select * from blog_tag where tag_name=?", t.TagName); e != nil {
		return false
	}
	return true
}

func (t *Tag) GetOne() (Tag, error) {
	var tag Tag
	if e := db.Get(&tag, "select * from blog_tag where id=?", t.ID); e != nil {
		return Tag{}, e
	}
	return tag, nil
}

func GetAllTags() ([]Tag, error) {
	tags := make([]Tag, 0)
	if err := db.Select(&tags, "select * from blog_tag"); err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *Tag) Add() (Tag, error) {
	r, e := db.Exec("insert into blog_tag (tag_name) values (?)", t.TagName)
	if e != nil {
		return Tag{}, e
	}
	id, _ := r.LastInsertId()
	return Tag{int(id), t.TagName}, nil
}

func (t *Tag) Delete() error {
	_, e := db.Exec("delete from blog_tag where id=?", t.ID)
	if e != nil {
		return e
	}
	return nil
}

func (t *Tag) Edit() error {
	_, e := db.Exec("update blog_tag set tag_name=? where id=?", t.TagName, t.ID)
	if e != nil {
		return e
	}
	return nil
}

func GetTagIDByName(name string) (int, error) {
	var id int
	if e := db.Get(&id, "select id from blog_tag where tag_name=?", name); e != nil {
		return id, e
	}
	return id, nil
}

func GetTagIDByArticle(articleID int) ([]int, error) {

	var tagIDs []int
	if e := db.Select(&tagIDs, "SELECT ta.tag_id FROM blog_article a, blog_tag_article ta WHERE a.id=ta.article_id AND a.id=?", articleID); e != nil {
		return nil, e
	}
	return tagIDs, nil
}

// 所有category的api方法
func (c *Category) IfExistByName() bool {
	category := new(Category)
	if e := db.Get(category, "select * from blog_category where category_name=?", c.CategoryName); e != nil {
		return false
	}
	return true
}

func (c *Category) GetOne() (Category, error) {
	var category Category
	if e := db.Get(&category, "select * from blog_category where id=?", c.ID); e != nil {
		return Category{}, e
	}
	return category, nil
}

func GetAllCategories() ([]Category, error) {
	categories := make([]Category, 0)
	if err := db.Select(&categories, "select * from blog_category"); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) Add() (Category, error) {
	r, e := db.Exec("insert into blog_category (category_name) values (?)", c.CategoryName)
	if e != nil {
		return Category{}, e
	}
	id, _ := r.LastInsertId()
	return Category{int(id), c.CategoryName}, nil
}

func (c *Category) Delete() error {
	_, e := db.Exec("delete from blog_category where id=?", c.ID)
	if e != nil {
		return e
	}
	return nil
}

func (c *Category) Edit() error {
	_, e := db.Exec("update blog_category set category_name=? where id=?", c.CategoryName, c.ID)
	if e != nil {
		return e
	}
	return nil
}

// 所有article的api方法
func (a *Article) IfExistByName() bool {
	var articleID int
	if e := db.Get(&articleID, "select id from blog_article where title=? and id<>?", a.Title, a.ID); e != nil {
		return false
	}
	return true
}

func (a *Article) GetOne() (Article, error) {
	var article Article
	if e := db.Get(&article, "select * from blog_article where id=?", a.ID); e != nil {
		return Article{}, e
	}
	tagIDs, e := GetTagIDByArticle(a.ID)
	if e != nil {
		return Article{}, e
	}
	article.TagID = tagIDs
	return article, nil
}

func (a *Article) Add() (Article, error) {
	createTime := time.Now().Format(utils.AppInfo.TimeFormat)
	tx := db.MustBegin()

	r, e := tx.Exec("insert into blog_article (title,content,category_id,created_time) values (?,?,?,?)", a.Title, a.Content, a.CategoryID, createTime)
	if e != nil {
		HandleRollback(tx)
		return Article{}, e
	}
	articleID, _ := r.LastInsertId()
	if len(a.TagID) > 0 {
		for _, tagID := range a.TagID {
			_, e := tx.Exec("insert into blog_tag_article (tag_id, article_id) values (?, ?)", tagID, articleID)
			if e != nil {
				HandleRollback(tx)
				return Article{}, e
			}
		}
	}

	HandleCommit(tx)
	return Article{int(articleID), a.Title, a.Content, a.CategoryID, a.TagID, createTime, a.UpdatedTime}, nil
}

func (a *Article) Delete() error {
	tx := db.MustBegin()
	if _, e := tx.Exec("delete from blog_tag_article where article_id=?", a.ID); e != nil {
		HandleRollback(tx)
		return e
	}
	if _, e := tx.Exec("delete from blog_article where id=?", a.ID); e != nil {
		HandleRollback(tx)
		return e
	}
	HandleCommit(tx)
	return nil
}

func (a *Article) Edit() error {
	tx := db.MustBegin()

	if _, e := tx.Exec("update blog_article set title=?,content=?,category_id=?,updated_time=? where id=?", a.Title, a.Content, a.CategoryID, a.UpdatedTime, a.ID); e != nil {
		HandleRollback(tx)
		return e
	}

	if _, e := tx.Exec("delete from blog_tag_article where article_id=?", a.ID); e != nil {
		HandleRollback(tx)
		return e
	}

	if len(a.TagID) > 0 {
		for _, tagID := range a.TagID {
			_, e := tx.Exec("insert into blog_tag_article (tag_id, article_id) values (?, ?)", tagID, a.ID)
			if e != nil {
				HandleRollback(tx)
				return e
			}
		}
	}

	HandleCommit(tx)
	return nil
}

var pageSize = utils.AppInfo.PageSize

func genArticles(baseSql, key string, page int) ([]Article, int, error) {
	TempArticles := make([]Article, 0)
	selectSql := fmt.Sprintf(baseSql, key) + fmt.Sprintf(" limit %d offset %d", pageSize, (page-1)*pageSize)

	if err := db.Select(&TempArticles, selectSql); err != nil {
		return nil, 0, err
	}
	total := getCount(baseSql)

	articles := make([]Article, 0)
	if len(TempArticles) > 0 {
		for _, value := range TempArticles {
			tagIDs, e := GetTagIDByArticle(value.ID)
			if e != nil {
				return nil, 0, e
			}
			value.TagID = tagIDs
			articles = append(articles, value)
		}
	}
	return articles, total, nil
}

func GetArticles(page int) ([]Article, int, error) {
	baseSql := "select %s from blog_article"
	articles, total, e := genArticles(baseSql, "*", page)
	if e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}

func getCount(sql string) int {
	var count int
	if e := db.Get(&count, fmt.Sprintf(sql, "count(1)")); e != nil {
		return 0
	}
	return count
}

func GetArticlesByCategory(name string, page int) ([]Article, int, error) {
	baseSql := "SELECT %s FROM blog_article a  INNER JOIN blog_category c ON a.category_id=c.id AND c.category_name='" + name + "'"
	articles, total, e := genArticles(baseSql, "a.*", page)
	if e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}

func GetArticlesByTag(name string, page int) ([]Article, int, error) {
	tagID, e := GetTagIDByName(name)
	if e != nil {
		return nil, 0, e
	}
	id := strconv.Itoa(tagID)
	baseSql := "SELECT %s FROM blog_article a  INNER JOIN blog_tag_article ta ON a.id=ta.article_id AND ta.tag_id='" + id + "'"

	articles, total, e := genArticles(baseSql, "a.*", page)
	if e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}
