package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticleComments(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Article{ID: id}

	comments, e := r.GetCommentsByArticle()
	if e != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40020, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, comments, nil))
	return
}

func CreateComment(c *gin.Context) {
	comment := &service.Comment{}
	if e := c.ShouldBindJSON(comment); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40024, nil, e))
		return
	}

	co, e := comment.Create()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40024, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, co, nil))
	return
}

func GetAllComments(c *gin.Context) {
	limit := c.DefaultQuery("limit", "")
	page := c.DefaultQuery("page", "")
	title := c.DefaultQuery("title", "")

	if title != "" {
		comments, err := service.GetCommentsByArticleName(title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40032, nil, err))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, comments, nil))
		return
	}

	comments, e := service.Comment{}.GetAll(limit, page)
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40032, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, comments, nil))
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Comment{ID: uint(id)}

	comment, i := r.GetOne()
	if i != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40006, nil, i))
		return
	}
	if e := comment.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40011, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, comment, nil))
}

func GetRencentComments(c *gin.Context) {
	comment := service.Comment{}
	data, e := comment.GetRecentComments()
	if e != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40020, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
	return
}
