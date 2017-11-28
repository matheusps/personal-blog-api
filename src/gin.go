package main

import (
	"github.com/gin-gonic/gin"
)

func provideServer() *gin.Engine {

	app := gin.Default()
	app.Use(CORSMiddleware())
	app.GET("/posts/", GetPosts)
	app.GET("/posts/:id", GetPost)
	app.POST("/posts", CreatePost)
	app.PUT("posts/:id", UpdatePosts)
	app.DELETE("/post/:id", DeletePosts)

	return app
}
