package routes

import (
	"myblogapi/controllers"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(r *gin.Engine) {
	blogController := controllers.BlogController{}
	r.GET("/blogs", blogController.GetAllBlogs)
	r.POST("/blogs", blogController.PostBlog)

}
