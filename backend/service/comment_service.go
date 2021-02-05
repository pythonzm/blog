package service

import (
	"backend/tools"
	"backend/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Comment struct {
	ID          uint          `json:"id" db:"id"`
	UserName    string        `json:"username" db:"username" binding:"required,max=16"`
	IsAuthor    bool          `json:"is_author" db:"is_author"`
	ParentID    sql.NullInt64 `json:"parent_id" db:"parent_id"` // 回复某条评论的ID
	RootID      sql.NullInt64 `json:"root_id" db:"root_id"`     // 根评论ID
	ArticleID   int           `json:"article_id" db:"article_id" binding:"required"`
	Content     string        `json:"content" db:"content" binding:"required"`
	CreatedTime string        `json:"created_time" db:"created_time"`
}

type RootComment struct {
	Comment
	Children []Comment `json:"children"`
}

type ArticleComments struct {
	Items []RootComment `json:"items"`
	Total uint          `json:"total"`
}

type CommentAndArticle struct {
	C Comment `json:"comment"`
	A Article `json:"article"`
}

type AllComments struct {
	Items []CommentAndArticle `json:"items"`
	Total uint                `json:"total"`
}

func (c *Comment) Create() (Comment, error) {
	var isAuthor int
	createTime := time.Now().Format(utils.AppInfo.TimeFormat)
	if c.IsAuthor {
		isAuthor = 1
	} else {
		isAuthor = 0
		go func(articleId int) {
			if utils.MailInfo.Enable {
				article, _ := Article{ID: int(articleId)}.GetOne(SetAdmin("true"))
				title := article.A.Title
				if e := utils.MailNotice(title, articleId); e != nil {
					utils.WriteErrorLog(fmt.Sprintf("[ %s ] 评论发送邮件通知失败, %v\n", time.Now().Format(utils.AppInfo.TimeFormat), e))
				}
			}
		}(c.ArticleID)
	}
	comment, e := db.Exec("insert into blog_comment (username,is_author,parent_id,root_id,article_id,content, created_time) values (?,?,?,?,?,?,?)", c.UserName, isAuthor, c.ParentID, c.RootID, c.ArticleID, c.Content, createTime)
	if e != nil {
		return Comment{}, e
	}
	commentID, _ := comment.LastInsertId()
	return Comment{uint(commentID), c.UserName, c.IsAuthor, c.ParentID, c.RootID, c.ArticleID, c.Content, createTime}, nil
}

func (c *Comment) Delete() error {
	_, e := db.Exec("delete from blog_comment where id=?", c.ID)
	return e
}

func (c *Comment) GetOne() (Comment, error) {
	var comment Comment
	e := db.Get(&comment, "select * from blog_comment where id=?", c.ID)
	return comment, e
}

func (c Comment) GetAll(limit, page string) (data AllComments, err error) {
	baseSql := "select %s from blog_comment s"
	var p, l int
	if limit != "" && page != "" {
		p, _ = strconv.Atoi(page)
		l, _ = strconv.Atoi(limit)
	}
	key := commentCacheKey(p, l)

	cacheData, e := getCommentCache(key)
	if e != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 读取缓存失败, %v\n", time.Now().Format(utils.AppInfo.TimeFormat), e))
	}
	if cacheData.Total != 0 {
		return cacheData, nil
	}

	var comments []Comment
	var allComments []CommentAndArticle
	offset := (p - 1) * l
	selectSql := fmt.Sprintf(baseSql, "s.*") + fmt.Sprintf(" ORDER BY s.id DESC limit %d offset %d", l, offset)
	if err = db.Select(&comments, selectSql); err != nil {
		return
	}

	for _, value := range comments {
		article, _ := Article{ID: int(value.ArticleID)}.GetArticle()
		allComments = append(allComments, CommentAndArticle{value, article})
	}

	var total uint
	if err = db.Get(&total, fmt.Sprintf(baseSql, "count(id)")); err != nil {
		return
	}

	data.Total = total
	data.Items = allComments

	if e := setCommentCache(key, data); e != nil {
		utils.WriteErrorLog(fmt.Sprintf("[ %s ] 写入缓存失败, %v\n", time.Now().Format(utils.AppInfo.TimeFormat), e))
	}

	return
}

func GetCommentsByArticleName(articleName string) (data AllComments, err error) {
	var comments []Comment
	var allComments []CommentAndArticle

	err = db.Select(&comments, "select c.* FROM blog_comment c INNER JOIN blog_article a ON c.article_id=a.id WHERE a.title=?", articleName)
	if err != nil {
		return
	}
	for _, value := range comments {
		allComments = append(allComments, CommentAndArticle{value, Article{ID: int(value.ArticleID), Title: articleName}})
	}
	data = AllComments{allComments, uint(len(comments))}
	return
}

func (a Article) GetCommentsByArticle() (ArticleComments, error) {
	var rootComments []RootComment
	roots, e := a.GetRootCommentsByArticle()
	if e != nil {
		return ArticleComments{}, e
	}
	for _, value := range roots {
		comments, e := value.GetChildren()
		if e != nil {
			return ArticleComments{}, e
		}
		rootComments = append(rootComments, RootComment{value, comments})
	}
	count, _ := a.getCommentsCount()
	return ArticleComments{rootComments, count}, nil
}

func (a Article) GetRootCommentsByArticle() ([]Comment, error) {
	comments := make([]Comment, 0)
	err := db.Select(&comments, "select * from blog_comment where article_id=? and root_id is NULL", a.ID)
	return comments, err
}

func (c *Comment) GetChildren() ([]Comment, error) {
	comments := make([]Comment, 0)
	err := db.Select(&comments, "select * from blog_comment where root_id=?", c.ID)
	return comments, err
}

func (a Article) getCommentsCount() (uint, error) {
	var count uint
	err := db.Get(&count, "select count(id) from blog_comment where article_id=?", a.ID)
	return count, err
}

func commentCacheKey(limit, page int) string {
	if limit == 0 || page == 0 {
		return fmt.Sprintf("comment_%d_%d", 10, 1)
	}
	return fmt.Sprintf("comment_%d_%d", limit, page)
}

func getCommentCache(key string) (s AllComments, err error) {
	data, e := tools.GetKey(key)
	if e != nil || data == nil {
		return s, e
	}

	v, ok := data.([]uint8)
	if ok {
		if e := json.Unmarshal([]byte(v[:]), &s); e != nil {
			return s, e
		}
		return s, nil
	} else {
		return s, errors.New("返回数据类型有误，json无法解析")
	}
}

func setCommentCache(key string, value AllComments) error {
	marshal, _ := json.Marshal(value)
	e := tools.SetKey(key, marshal, tools.SetTimeout(true))
	return e
}

func (c Comment) GetRecentComments() (data []map[string]string, err error) {
	var comments []Comment
	if err = db.Select(&comments, "SELECT * FROM blog_comment ORDER BY id DESC LIMIT 10"); err != nil {
		return
	}
	for _, comment := range comments {
		a := Article{ID: comment.ArticleID}
		title, _ := a.GetTitleByID()
		data = append(data, map[string]string{"title": title, "content": comment.Content, "create_time": comment.CreatedTime})
	}
	return
}
