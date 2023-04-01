package main

import (
	"challenge/project1/app"
	"challenge/project1/config"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	err = config.InitPostgres()
	if err != nil {
		panic(err)
	}

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
