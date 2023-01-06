package controllers

import (
	"encoding/json"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(context *gin.Context) {
	var comment models.Comment
	//var user models.User
	//user1 := database.Instance.Where("id = ?", 1).First(&user)

	if err := context.ShouldBindJSON(&comment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&comment)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"commentId": comment.ID, "movieId": comment.MovieID, "userId": comment.UserID, "text": comment.Text})
}

func DeleteComment(context *gin.Context) {
	id := context.Param("commentId")

	var comment models.Comment
	log.Println(id)
	record := database.Instance.Where("id = ?", id).First(&comment)
	if record.Error == nil {
		context.JSON(200, gin.H{"success": "Movie #" + id + " deleted"})
		record.Delete(&comment)
		context.Abort()
		return
	} else {
		context.JSON(404, gin.H{"error": "User not found"})
	}
}

func UpdateComment(context *gin.Context) {

	id := context.Param("commentId")
	var comment models.Comment
	var newComment models.Comment

	if err := json.NewDecoder(context.Request.Body).Decode(&newComment); err != nil {
		context.JSON(404, gin.H{"error": "could not parse json"})
		return
	}

	record := database.Instance.Where("id = ?", id).First(&comment)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "Movie not found"})
		context.Abort()
		return
	}
	record.Updates(&newComment)
	context.JSON(200, gin.H{"success": "Movie #" + id + " Updated"})
	context.Abort()
	return
}

/*func GetMovie(context *gin.Context) {

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
}*/

func GetUserCommentsByMovie(context *gin.Context) {
	var comment []models.Comment

	id := context.Param("movieId")
	records := database.Instance.Where("movie_id = ?", id).First(&comment)

	//records := database.Instance.Find(&comments)
	if records.Error != nil {
		context.JSON(404, gin.H{"error": "for this movie Comment not found"})
		context.Abort()
		return
	}
	context.JSON(200, &comment)
	context.Abort()
	return
}
func GetUserCommentsByUser(context *gin.Context) {
	var comment []models.Comment

	id := context.Param("userId")
	records := database.Instance.Where("user_id = ?", id).First(&comment)

	// records := database.Instance.Find(&comment)
	if records.Error != nil {
		context.JSON(404, gin.H{"error": "for this User Comment not found"})
		context.Abort()
		return
	}
	context.JSON(200, &comment)
	context.Abort()
	return
}
