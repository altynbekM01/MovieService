package database

import (
	"jwt-authentication-golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Profile{})
	Instance.AutoMigrate(&models.Bookmark{})
	Instance.AutoMigrate(&models.Comment{})
	Instance.AutoMigrate(&models.Genre{})
	Instance.AutoMigrate(&models.Celebrity{})
	Instance.AutoMigrate(&models.Movie{})

	// Instance.AutoMigrate(&models.Movie{})
	log.Println("Database Migration Completed!")
}
