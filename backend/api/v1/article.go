package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	category := c.DefaultQuery("category", "")
	tag := c.DefaultQuery("tag", "")
	key := c.DefaultQuery("key", "")
	status := c.DefaultQuery("status", "")
	admin := c.DefaultQuery("admin", "")

	if category != "" {
		data, e := service.GetArticlesByCategory(category, service.SetLimitPageAdmin(limit, page, admin), service.SetCategory(category))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	if tag != "" {
		data, e := service.GetArticlesByTag(tag, service.SetLimitPageAdmin(limit, page, admin), service.SetTag(tag))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	if key != "" || status != "" {
		data, e := service.SearchArticle(key, status, service.SetLimitPageAdmin(limit, page, admin), service.SetSearch(true))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	a := service.Article{}
	data, e := a.GetAll(service.SetLimitPageAdmin(limit, page, admin))
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
	return
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Article{ID: id}

	articleDetail, e := r.GetOne()
	if e != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40020, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, articleDetail, nil))
	return
}

func CreateArticle(c *gin.Context) {
	article := &service.Article{}
	if e := c.ShouldBindJSON(article); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40024, nil, e))
		return
	}

	a, e := article.Create()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40024, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, a, nil))
	return
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Article{ID: id}

	articleDetail, _ := r.GetOne()
	article := articleDetail.A

	if e := r.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40026, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, article, nil))
	return
}

func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Article{ID: id}

	articleDetail, _ := r.GetOne()
	article := articleDetail.A

	if e := c.ShouldBindJSON(&article); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40024, nil, e))
		return
	}

	if e := article.Edit(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40025, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, article, nil))
	return
}
