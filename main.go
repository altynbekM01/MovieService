package main

import (
	"fmt"
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"
	"jwt-authentication-golang/models"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
)

func main() {
	URL := fmt.Sprintf("root:@tcp(localhost:3306)/golang2?parseTime=true")
	fmt.Println(URL)
	database.Connect(URL)
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

// asem push
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)

		api.POST("/movie/add", controllers.CreateMovie)
		api.DELETE("/movie/delete/:movieId", controllers.DeleteMovie)
		api.PUT("/movie/update/:movieId", controllers.UpdateMovie)

		api.POST("/celebrity/add", controllers.CreateCelebrity)
		api.DELETE("/celebrity/delete/:celebrityId", controllers.DeleteCelebrity)
		api.PUT("/celebrity/update/:celebrityId", controllers.UpdateCelebrity)

		api.POST("/comment/add", controllers.CreateComment)
		api.DELETE("/comment/delete/:commentId", controllers.DeleteComment)
		api.PUT("/comment/update/:commentId", controllers.UpdateComment)

		api.POST("/genre/add", controllers.CreateGenre)
		api.DELETE("/genre/delete/:genreId", controllers.DeleteGenre)
		api.PUT("/genre/update/:genreId", controllers.UpdateGenre)
		api.GET("/genre/get/:genreId", controllers.GetGenre)
		api.GET("/genres", controllers.GetGenres)

		api.GET("/getCommentsOfMovie/:movieId", controllers.GetUserCommentsByMovie)
		api.GET("/getCommentsOfUser/:userId", controllers.GetUserCommentsByUser)

		api.GET("/celebrity/get/:celebrityId", controllers.GetCelebrity)
		api.GET("/celebrities", controllers.GetCelebrites)

		api.GET("/movie/search", func(context *gin.Context) {
			var movies []models.Movie

			sql := "SELECT * FROM movies"

			if s := context.Query("s"); s != "" {
				sql = fmt.Sprintf("%s WHERE name LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
			}

			if sort := context.Query("sort"); sort != "" {
				sql = fmt.Sprintf("%s ORDER BY rating %s", sql, sort)
			}

			database.Instance.Raw(sql).Scan(&movies)
			//context.JSON(200, &movies)

			page, _ := strconv.Atoi(context.Query("page"))
			perPage := 3
			var total int64

			database.Instance.Raw(sql).Count(&total)

			sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

			database.Instance.Raw(sql).Scan(&movies)

			context.JSON(200, fiber.Map{
				"data":      &movies,
				"total":     total,
				"page":      page,
				"last_page": math.Ceil(float64(total / int64(perPage))),
			})

			return
		})

		api.GET("/movie/get/:movieId", controllers.GetMovie)
		// crud_movie := api.Group(("/movie")).Use()
		// {
		// 	crud_movie.POST("/movie/add", controllers.CreateMovie)
		// 	crud_movie.DELETE("/movie/delete/:movieId", controllers.DeleteMovie)
		// 	crud_movie.PUT("/movie/update/:movieId", controllers.UpdateMovie)
		// 	crud_movie.GET("/movie/get/:movieId", controllers.GetMovie)
		// }
		api.GET("/movies", controllers.GetMovies)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
