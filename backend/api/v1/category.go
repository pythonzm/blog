package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCategories(c *gin.Context) {
	categories, e := service.Category{}.GetAll()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40008, nil, e))
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, categories, nil))
}

func GetCategoryCount(c *gin.Context) {
	a := service.Category{}
	count, e := a.GetCategoryCount()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40035, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
	return
}

func CreateCategory(c *gin.Context) {
	category := &service.Category{}
	if e := c.ShouldBindJSON(category); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}
	ca, e := category.Create()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40009, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, ca, nil))
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Category{ID: id}

	category, i := r.GetOne()
	if i != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40006, nil, i))
		return
	}
	if e := category.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40011, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, category, nil))
}

func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ca := service.Category{ID: id}

	if e := c.ShouldBindJSON(&ca); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}

	if e := ca.Edit(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40010, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, ca, nil))
}
