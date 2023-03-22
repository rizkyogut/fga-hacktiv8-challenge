package routers

import (
	"challenge6/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/book/:bookID", controllers.GetBook)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeletedBook)

	return router
}
