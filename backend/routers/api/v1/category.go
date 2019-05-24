package v1

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCategories(c *gin.Context) {
	categories, e := models.GetAllCategories()
	if e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.CategoryGetAllError, e, nil)
	}
	utils.Response(c, http.StatusOK, utils.Success, nil, categories, len(categories))
}

func AddCategory(c *gin.Context) {
	category := &models.Category{}
	if e := c.ShouldBindJSON(category); e != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, e, nil)
		return
	}
	if b := category.IfExistByName(); b {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.CategoryExistError, nil, nil)
		return
	}
	t, e := category.Add()
	if e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.CategoryAddError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusCreated, utils.Success, nil, t)
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Category{ID: id}

	category, i := r.GetOne()
	if i != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.CategoryGetError, i, nil)
		return
	}
	if e := category.Delete(); e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.CategoryDeleteError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, category)
}

func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Category{ID: id}

	category, i := r.GetOne()
	if i != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.CategoryGetError, i, nil)
		return
	}

	if e := c.ShouldBindJSON(&category); e != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, e, nil)
		return
	}
	if b := category.IfExistByName(); b {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.CategoryExistError, nil, nil)
		return
	}
	if e := category.Edit(); e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.CategoryEditError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, category)
}
