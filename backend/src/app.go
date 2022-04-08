package main

import (
	app "dcardHw/src/api"
	"dcardHw/src/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := app.StartServer()
	model.Init()
	app.Run(":8080")
}
