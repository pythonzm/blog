package v1

import (
	"backend/models"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllTags(c *gin.Context) {
	tags, e := models.GetAllTags()
	if e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.TagGetAllError, e, nil)
	}
	utils.Response(c, http.StatusOK, utils.Success, nil, tags, len(tags))
}

func AddTag(c *gin.Context) {
	tag := &models.Tag{}
	if e := c.ShouldBindJSON(tag); e != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, e, nil)
		return
	}
	if b := tag.IfExistByName(); b {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.TagExistError, nil, nil)
		return
	}
	t, e := tag.Add()
	if e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.TagAddError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusCreated, utils.Success, nil, t)
}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Tag{ID: id}

	tag, i := r.GetOne()
	if i != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.TagGetError, i, nil)
		return
	}
	if e := tag.Delete(); e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.TagDeleteError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, tag)
}

func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := models.Tag{ID: id}

	t, i := r.GetOne()
	if i != nil {
		utils.DefaultResponse(c, http.StatusNotFound, utils.TagGetError, i, nil)
		return
	}

	if e := c.ShouldBindJSON(&t); e != nil {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.InvalidParams, e, nil)
		return
	}
	if b := t.IfExistByName(); b {
		utils.DefaultResponse(c, http.StatusBadRequest, utils.TagExistError, nil, nil)
		return
	}
	if e := t.Edit(); e != nil {
		utils.DefaultResponse(c, http.StatusInternalServerError, utils.TagEditError, e, nil)
		return
	}
	utils.DefaultResponse(c, http.StatusOK, utils.Success, nil, t)
}
