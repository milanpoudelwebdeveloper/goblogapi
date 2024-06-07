package controllers

import (
	"fmt"
	"myblogapi/db"
	"myblogapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type BlogController struct{}

func (ctrl BlogController) GetAllBlogs(c *gin.Context) {
	query := "SELECT * FROM blog"
	blogs := []models.Blog{}
	err := db.DB.Select(&blogs, query)
	if err != nil {
		fmt.Println("Error: Could not fetch the blogs.", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while fetching the blogs.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All blogs fetched successfully.",
		"blogs":   blogs,
	})
}

func (ctrl BlogController) PostBlog(c *gin.Context) {
	var blog models.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		fmt.Println("Error: Could not bind the JSON.", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON.",
		})
		return
	}
	blog.MetaTitle = slug.Make(blog.Title)
	query := "INSERT INTO blog (title, metatitle, content, coverimage, published, featured, writtenby) VALUES (:title, :metatitle, :content, :coverimage, :published, :featured, :writtenby)"
	_, err = db.DB.NamedExec(query, blog)
	if err != nil {
		fmt.Println("Error: Could not insert the blog.", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong while inserting the blog.",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog inserted successfully.",
	})

}
