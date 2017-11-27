package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		fmt.Println(err)
	}

	return db
}

func AutoMigrateTypes() {
	db := ConnectDB()
	defer db.Close()

	db.AutoMigrate(&Post{})
}
