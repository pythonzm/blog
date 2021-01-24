package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllTags(c *gin.Context) {
	tags, e := service.Tag{}.GetAll()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40008, nil, e))
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, tags, nil))
}

func GetTagCount(c *gin.Context) {
	a := service.Tag{}
	count, e := a.GetTagCount()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40036, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, count, nil))
	return
}

func CreateTag(c *gin.Context) {
	tag := &service.Tag{}
	if e := c.ShouldBindJSON(tag); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}
	t, e := tag.Create()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40009, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, t, nil))
}

func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Tag{ID: id}

	tag, i := r.GetOne()
	if i != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40006, nil, i))
		return
	}
	if e := tag.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40011, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, tag, nil))
}

func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	t := service.Tag{ID: id}

	if e := c.ShouldBindJSON(&t); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}

	if e := t.Edit(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40010, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, t, nil))
}
