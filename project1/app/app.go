package app

import (
	"challenge/project1/config"
	"challenge/project1/handler"
	"challenge/project1/repository"
	"challenge/project1/route"
	"challenge/project1/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB, config.GORM.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
