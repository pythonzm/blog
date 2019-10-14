package v1

import (
	"backend/service"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllSoups(c *gin.Context) {
	limit := c.DefaultQuery("limit", "")
	page := c.DefaultQuery("page", "")
	soups, e := service.Soup{}.GetAll(limit, page)
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40008, nil, e))
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, soups, nil))
}

func CreateSoup(c *gin.Context) {
	soup := &service.Soup{}
	if e := c.ShouldBindJSON(soup); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}
	s, e := soup.Create()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40009, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, s, nil))
}

func DeleteSoup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r := service.Soup{ID: id}

	soup, i := r.GetOne()
	if i != nil {
		c.JSON(http.StatusNotFound, utils.GenResponse(40006, nil, i))
		return
	}
	if e := soup.Delete(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40011, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, soup, nil))
}

func EditSoup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	s := service.Soup{ID: id}

	if e := c.ShouldBindJSON(&s); e != nil {
		c.JSON(http.StatusBadRequest, utils.GenResponse(40000, nil, e))
		return
	}

	if e := s.Edit(); e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40010, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, s, nil))
}

func GetRandSoup(c *gin.Context) {
	soup, e := service.Soup{}.GetRandOne()
	if e != nil {
		c.JSON(http.StatusInternalServerError, utils.GenResponse(40008, nil, e))
		return
	}
	c.JSON(http.StatusOK, utils.GenResponse(20000, soup, nil))
}
