package app

import (
	"challenge/challenge7/config"
	"challenge/challenge7/handler"
	"challenge/challenge7/repository"
	"challenge/challenge7/route"
	"challenge/challenge7/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
