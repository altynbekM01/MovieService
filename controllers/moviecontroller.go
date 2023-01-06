package controllers

import (
	"encoding/json"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMovie(context *gin.Context) {
	var movie models.Movie
	if err := context.ShouldBindJSON(&movie); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	log.Println(movie)
	// movie.Actor =
	record := database.Instance.Create(&movie)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"movieId": movie.ID, "Name": movie.Name, "Genre": movie.Genre, "Description": movie.Description})
}

func DeleteMovie(context *gin.Context) {
	id := context.Param("movieId")

	var movie models.Movie
	log.Println(id)
	record := database.Instance.Where("id = ?", id).First(&movie)
	if record.Error == nil {
		context.JSON(200, gin.H{"success": "Movie #" + id + " deleted"})
		record.Delete(&movie)
		context.Abort()
		return
	} else {
		context.JSON(404, gin.H{"error": "User not found"})
	}
}

func UpdateMovie(context *gin.Context) {

	id := context.Param("movieId")
	var movie models.Movie
	var newMovie models.Movie

	if err := json.NewDecoder(context.Request.Body).Decode(&newMovie); err != nil {
		context.JSON(404, gin.H{"error": "could not parse json"})
		return
	}

	record := database.Instance.Where("id = ?", id).First(&movie)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "Movie not found"})
		context.Abort()
		return
	}
	record.Updates(&newMovie)
	context.JSON(200, gin.H{"success": "Movie #" + id + " Updated"})
	context.Abort()
	return
}

func GetMovie(context *gin.Context) {

	id := context.Param("movieId")
	var movie models.Movie
	record := database.Instance.Where("id = ?", id).First(&movie)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "Movie not found"})
		context.Abort()
		return
	}
	// context.JSON(200, gin.H{"id": id, "Name": movie.Name, "Genre": movie.Genre, "Description": movie.Description})
	context.JSON(200, &movie)
	context.Abort()
	return
}

func GetMovies(context *gin.Context) {

	var movies []models.Movie
	// var genres []models.Genre
	// records := database.Instance.Find(&movies)
	records := database.Instance.Preload("Genre").Preload("Actor").Find(&movies)
	if records.Error != nil {
		context.JSON(404, gin.H{"error": "Movie not found"})
		context.Abort()
		return
	}
	// records := database.Instance.Model(&movies).Association("Genres").Find(&genres)
	context.JSON(200, &movies)
	context.Abort()
	return
}
