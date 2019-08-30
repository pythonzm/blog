package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetArticles(c *gin.Context) {
	category := c.Query("category")
	tag := c.Query("tag")
	key := c.Query("key")
	status := c.Query("status")
	limit := c.Query("limit")
	page := c.Query("page")

	if category != "" {
		data, e := service.GetArticlesByCategory(category, service.SetLimitPage(limit, page))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	if tag != "" {
		data, e := service.GetArticlesByTag(tag, service.SetLimitPage(limit, page))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	if key != "" || status != "" {
		data, e := service.SearchArticle(key, status, service.SetLimitPage(limit, page), service.SetSearch(true))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	a := service.Article{}
	data, e := a.GetAll(service.SetLimitPage(limit, page))
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

	article, e := r.GetOne()
	if e != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40020, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, article, nil))
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

	article, _ := r.GetOne()

	if e := article.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40026, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, article, nil))
	return
}

func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Article{ID: id}

	article, _ := r.GetOne()

	if e := c.ShouldBindJSON(&article); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40024, nil, e))
		return
	}
	article.UpdatedTime = time.Now().Format(utils.AppInfo.TimeFormat)

	if e := article.Edit(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40025, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, article, nil))
	return
}
