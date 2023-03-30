package route

import (
	"challenge/challenge7/handler"
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/book") // prefix
	{
		api.POST("", server.CreateBook) // /employees
		api.GET("/all", server.GetAllBook)
		api.GET("/:id", server.GetBookByID)   // /employee/:id
		api.PUT("/:id", server.UpdateBook)    // /employee/:id
		api.DELETE("/:id", server.DeleteBook) // /employee/:id
	}
}
