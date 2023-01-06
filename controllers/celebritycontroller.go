package controllers

import (
	"encoding/json"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCelebrity(context *gin.Context) {
	var celebrity models.Celebrity
	if err := context.ShouldBindJSON(&celebrity); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&celebrity)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"celebrityId": celebrity.ID, "firstName": celebrity.FirstName, "lastName": celebrity.LastName, "is_actor": celebrity.Is_actor, "is_producer": celebrity.Is_producer, "Movie": celebrity.Movie})
}

func DeleteCelebrity(context *gin.Context) {
	id := context.Param("celebrityId")

	var celebrity models.Celebrity
	log.Println(id)
	record := database.Instance.Where("id = ?", id).First(&celebrity)
	if record.Error == nil {
		context.JSON(200, gin.H{"success": "Celebrity #" + id + " deleted"})
		record.Delete(&celebrity)
		context.Abort()
		return
	} else {
		context.JSON(404, gin.H{"error": "Celebrity not found"})
	}
}

func UpdateCelebrity(context *gin.Context) {

	id := context.Param("celebrityId")
	var celebrity models.Celebrity
	var newCelebrity models.Celebrity

	if err := json.NewDecoder(context.Request.Body).Decode(&newCelebrity); err != nil {
		context.JSON(404, gin.H{"error": "could not parse json"})
		return
	}

	record := database.Instance.Where("id = ?", id).First(&celebrity)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "Celebrity not found"})
		context.Abort()
		return
	}
	record.Updates(&newCelebrity)
	context.JSON(200, gin.H{"success": "Celebrity #" + id + " Updated"})
	context.Abort()
	return
}

func GetCelebrity(context *gin.Context) {

	id := context.Param("celebrityId")
	var celebrity models.Celebrity
	record := database.Instance.Where("id = ?", id).First(&celebrity)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "Celebrity not found"})
		context.Abort()
		return
	}
	// context.JSON(200, gin.H{"id": id, "Name": movie.Name, "Genre": movie.Genre, "Description": movie.Description})
	context.JSON(200, &celebrity)
	context.Abort()
	return
}

func GetCelebrites(context *gin.Context) {

	var celebrity []models.Celebrity
	records := database.Instance.Find(&celebrity)
	if records.Error != nil {
		context.JSON(404, gin.H{"error": "Celebrity not found"})
		context.Abort()
		return
	}
	context.JSON(200, &celebrity)
	context.Abort()
	return
}
