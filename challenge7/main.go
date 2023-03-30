package main

import (
	"challenge/challenge7/app"
	"challenge/challenge7/config"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	err = config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
