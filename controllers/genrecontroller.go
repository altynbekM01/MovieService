package controllers

import (
	"encoding/json"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGenre(context *gin.Context) {
	var genre models.Genre
	if err := context.ShouldBindJSON(&genre); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&genre)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"genreId": genre.ID, "Name": genre.Name, "Movies": genre.Movie})
}

func DeleteGenre(context *gin.Context) {
	id := context.Param("genreId")

	var genre models.Genre
	log.Println(id)
	record := database.Instance.Where("id = ?", id).First(&genre)
	if record.Error == nil {
		context.JSON(200, gin.H{"success": "genre #" + id + " deleted"})
		record.Delete(&genre)
		context.Abort()
		return
	} else {
		context.JSON(404, gin.H{"error": "User not found"})
	}
}

func UpdateGenre(context *gin.Context) {

	id := context.Param("genreId")
	var genre models.Genre
	var newGenre models.Genre

	if err := json.NewDecoder(context.Request.Body).Decode(&newGenre); err != nil {
		context.JSON(404, gin.H{"error": "could not parse json"})
		return
	}
	record := database.Instance.Where("id = ?", id).First(&genre)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "genre not found"})
		context.Abort()
		return
	}
	record.Updates(&newGenre)
	context.JSON(200, gin.H{"success": "genre #" + id + " Updated"})
	context.Abort()
	return
}

func GetGenre(context *gin.Context) {
	id := context.Param("genreId")
	var genre models.Genre
	record := database.Instance.Where("id = ?", id).First(&genre)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "genre not found"})
		context.Abort()
		return
	}
	context.JSON(200, &genre)
	context.Abort()
	return
}

func GetGenres(context *gin.Context) {

	var genres []models.Genre
	records := database.Instance.Find(&genres)
	if records.Error != nil {
		context.JSON(404, gin.H{"error": "genre not found"})
		context.Abort()
		return
	}
	context.JSON(200, &genres)
	context.Abort()
	return
}
