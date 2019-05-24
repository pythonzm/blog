package routers

import (
	"backend/middleware"
	"backend/routers/api"
	"backend/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(middleware.CustomLogger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWt())
	{
		apiv1.GET("/tags", v1.GetAllTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/categories", v1.GetAllCategories)
		apiv1.POST("/categories", v1.AddCategory)
		apiv1.PUT("/categories/:id", v1.EditCategory)
		apiv1.DELETE("/categories/:id", v1.DeleteCategory)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
	}

	return r
}
