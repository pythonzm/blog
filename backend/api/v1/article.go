package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArticles(c *gin.Context) {
	limit := c.DefaultQuery("limit", "")
	page := c.DefaultQuery("page", "")
	category := c.DefaultQuery("category", "")
	q := c.DefaultQuery("q", "")
	tag := c.DefaultQuery("tag", "")
	key := c.DefaultQuery("key", "")
	status := c.DefaultQuery("status", "")
	admin := c.DefaultQuery("admin", "")

	if category != "" {
		data, e := service.GetArticlesByCategory(service.SetLimitPage(limit, page), service.SetAdmin(admin), service.SetCategory(category))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	if tag != "" {
		data, e := service.GetArticlesByTag(service.SetLimitPage(limit, page), service.SetAdmin(admin), service.SetTag(tag))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}

	if key != "" || status != "" {
		data, e := service.SearchArticle(key, status, service.SetLimitPage(limit, page), service.SetAdmin(admin), service.SetSearch(true))
		if e != nil {
			c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
			return
		}
		c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
		return
	}
	if q != "" {
		var articles service.Articles
		var e error
		if utils.ESInfo.Enable {
			articles, e = service.SearchFromES(service.SetQ(q), service.SetLimitPage(limit, page))
			if e != nil {
				c.JSON(http.StatusInternalServerError, utils.GenResponse(40033, nil, e))
				return
			}
		} else {
			articles, e = service.SearchArticle(q, "published", service.SetLimitPage(limit, page), service.SetSearch(true))
			if e != nil {
				c.JSON(http.StatusInternalServerError, utils.GenResponse(40033, nil, e))
				return
			}
		}

		c.JSON(http.StatusOK, utils.GenResponse(20000, articles, nil))
		return
	}

	a := service.Article{}
	data, e := a.GetAll(service.SetLimitPage(limit, page), service.SetAdmin(admin))
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40022, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, data, nil))
	return
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	admin := c.DefaultQuery("admin", "")
	r := service.Article{ID: id}

	articleDetail, e := r.GetOne(service.SetAdmin(admin))
	if e != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40020, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, articleDetail, nil))
	return
}

func GetArticleCount(c *gin.Context) {
	a := service.Article{}
	count, e := a.GetArticleCount()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40034, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
	return
}

func GetArticleCountByCategory(c *gin.Context) {
	a := service.Article{}
	count, e := a.GetArticleCountByCategory()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40034, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
	return
}

func GetArticleCountByTag(c *gin.Context) {
	a := service.Article{}
	count, e := a.GetArticleCountByTag()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40034, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
	return
}

func GetArticleRank(c *gin.Context) {
	a := service.Article{}
	count, e := a.GetArticleRank()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40039, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
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
	admin := c.DefaultQuery("admin", "")
	r := service.Article{ID: id}

	articleDetail, _ := r.GetOne(service.SetAdmin(admin))
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
	admin := c.DefaultQuery("admin", "")
	r := service.Article{ID: id}

	articleDetail, _ := r.GetOne(service.SetAdmin(admin))
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
