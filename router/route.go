package router

import (
	// "net/http"
	"project/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	// http.HandleFunc("/books", controller.GetBooks)
	// http.HandleFunc("/books/", controller.GetBookByID)
	// http.HandleFunc("/books", controller.CreateBook)
	// http.HandleFunc("/books/", controller.UpdateBook)
	// http.HandleFunc("/books/", controller.DeleteBook)

	r := gin.Default()

	r.GET("/books", controller.GetBooks)
	r.GET("/books/:id", controller.GetBookByID)
	r.POST("/books", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)

	return r
}
