package v1

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetArticles(c *gin.Context) {
	category := c.Query("category")
	tag := c.Query("tag")
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, err, nil)
		return
	}
	if category != "" {
		articles, total, e := models.GetArticlesByCategory(category, p)
		if e != nil {
			utils.DefaultResponse(c, http.StatusInternalServerError, utils.ArticleGetError, e, nil)
			return
		}
		utils.Response(c, http.StatusOK, utils.Success, nil, articles, total)
		return
	} else if tag != "" {
		articles, total, e := models.GetArticlesByTag(tag, p)
		if e != nil {
			utils.DefaultResponse(c, http.StatusInternalServerError, utils.ArticleGetError, e, nil)
			return
		}
		utils.Response(c, http.StatusOK, utils.Success, nil, articles, total)
		return
	}
	articles, total, e := models.GetArticles(p)
	if e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.ArticleGetAllError, e, nil)
		return
	}
	utils.Response(c, http.StatusOK, utils.Success, nil, articles, total)
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Article{ID: id}

	article, e := r.GetOne()
	if e != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.ArticleGetError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, article)
}

func AddArticle(c *gin.Context) {
	article := &models.Article{}
	if e := c.ShouldBindJSON(article); e != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, e, nil)
		return
	}
	if b := article.IfExistByName(); b {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.ArticleExistError, nil, nil)
		return
	}
	t, e := article.Add()
	if e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.ArticleAddError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusCreated, utils.Success, nil, t)
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Article{ID: id}

	article, i := r.GetOne()
	if i != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.ArticleGetError, i, nil)
		return
	}
	if e := article.Delete(); e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.ArticleDeleteError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, article)
}

func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Article{ID: id}

	article, i := r.GetOne()
	if i != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.ArticleGetError, i, nil)
		return
	}

	if e := c.ShouldBindJSON(&article); e != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, e, nil)
		return
	}
	article.UpdatedTime = time.Now().Format(utils.AppInfo.TimeFormat)

	if b := article.IfExistByName(); b {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.ArticleExistError, nil, nil)
		return
	}
	if e := article.Edit(); e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.ArticleEditError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, article)
}
