package service

import (
	"backend/models"
	"backend/utils"
	"fmt"
	"time"
)

type Article struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required,max=32"`
	Content     string `json:"content" db:"content" binding:"required"`
	CategoryID  int    `json:"category_id" db:"category_id" binding:"required"`
	TagID       []int  `json:"tag_id" binding:"required"`
	CreatedTime string `json:"created_time" db:"created_time"`
	UpdatedTime string `json:"updated_time" db:"updated_time"`
}

var db = models.DB

func (a *Article) Create() (Article, error) {
	createTime := time.Now().Format(utils.AppInfo.TimeFormat)

	r, e := db.Exec("insert into blog_article (title,content,category_id,created_time) values (?,?,?,?)", a.Title, a.Content, a.CategoryID, createTime)
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
	return Article{int(articleID), a.Title, a.Content, a.CategoryID, a.TagID, createTime, a.UpdatedTime}, nil
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
	if _, e := db.Exec("update blog_article set title=?,content=?,category_id=?,updated_time=? where id=?", a.Title, a.Content, a.CategoryID, a.UpdatedTime, a.ID); e != nil {
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

func (a Article) GetOne() (Article, error) {
	var one Article
	if e := db.Get(&one, "select * from blog_article where id=?", a.ID); e != nil {
		return Article{}, e
	}
	tagIDs, e := GetTagIDByArticle(a.ID)
	if e != nil {
		return Article{}, e
	}
	one.TagID = tagIDs
	return one, nil
}

func (a Article) GetAll(limit, page int) ([]Article, int, error) {
	baseSql := "select %s from blog_article a"
	articles, total, e := genArticles(baseSql, limit, page)
	if e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}

func GetArticlesByCategory(name string, limit, page int) ([]Article, int, error) {
	baseSql := "SELECT %s FROM blog_article a INNER JOIN blog_category c ON a.category_id=c.id AND c.category_name=" + name + ""
	articles, total, e := genArticles(baseSql, limit, page)
	if e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}

func GetArticlesByTag(name string, limit, page int) ([]Article, int, error) {
	baseSql := "SELECT %s FROM blog_article a  INNER JOIN blog_tag_article ta ON a.id=ta.article_id INNER JOIN blog_tag t ON ta.tag_id=t.id WHERE t.tag_name=" + name + ""
	articles, total, e := genArticles(baseSql, limit, page)
	if e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}

func genArticles(baseSql string, limit, page int) ([]Article, int, error) {
	articles := make([]Article, 0)
	var total int
	selectSql := fmt.Sprintf(baseSql, "a.*") + fmt.Sprintf(" limit %d offset %d", limit, (page-1)*limit)
	if err := db.Select(&articles, selectSql); err != nil {
		return nil, 0, err
	}

	if e := updateTagIDs(articles); e != nil {
		return nil, 0, e
	}

	if e := db.Get(&total, fmt.Sprintf(baseSql, "count(1)")); e != nil {
		return nil, 0, e
	}
	return articles, total, nil
}

func GetTagIDByArticle(articleID int) ([]int, error) {
	var tagIDs []int
	if e := db.Select(&tagIDs, "SELECT ta.tag_id FROM blog_article a, blog_tag_article ta WHERE a.id=ta.article_id AND a.id=?", articleID); e != nil {
		return nil, e
	}
	return tagIDs, nil
}

func updateTagIDs(articles []Article) error {
	if len(articles) > 0 {
		for key, value := range articles {
			tagIDs, e := GetTagIDByArticle(value.ID)
			if e != nil {
				return e
			}
			articles[key].TagID = tagIDs
		}
	}
	return nil
}
