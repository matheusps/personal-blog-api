package main

import (
	"github.com/gin-gonic/gin"
)

func provideServer() *gin.Engine {

	routes := gin.Default()
	routes.GET("/posts/", GetPosts)
	routes.GET("/posts/:id", GetPost)
	routes.POST("/posts", CreatePost)
	routes.PUT("posts/:id", UpdatePosts)
	routes.DELETE("/post/:id", DeletePosts)

	return routes
}
