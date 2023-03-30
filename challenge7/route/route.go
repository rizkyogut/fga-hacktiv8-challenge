package route

import (
	"challenge/challenge7/handler"
	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/book") // prefix
	{
		api.POST("", server.CreateBook)
		api.GET("/all", server.GetAllBook)
		api.GET("/:id", server.GetBookByID)
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)
	}
}
