package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func CreatePost(context *gin.Context) {

	db := ConnectDB()
	defer db.Close()

	var post Post
	context.BindJSON(&post)
	db.Create(&post)
	context.JSON(http.StatusOK, post)
}

func GetPosts(context *gin.Context) {

	db := ConnectDB()
	defer db.Close()

	var posts []Post

	if err := db.Find(&posts).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		context.JSON(http.StatusOK, posts)
	}
}

func GetPost(context *gin.Context) {

	db := ConnectDB()
	defer db.Close()

	id := context.Params.ByName("id")
	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		context.JSON(http.StatusOK, post)
	}
}

func UpdatePosts(context *gin.Context) {

	db := ConnectDB()
	defer db.Close()

	var post Post
	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	context.BindJSON(&post)
	db.Save(&post)
	context.JSON(http.StatusOK, post)
}

func DeletePosts(context *gin.Context) {

	db := ConnectDB()
	defer db.Close()

	id := context.Params.ByName("id")
	var post Post
	res := db.Where("id = ?", id).Delete(&post)
	fmt.Println(res)
	context.JSON(http.StatusOK, gin.H{
		"id #" + id: "deleted",
	})
}
