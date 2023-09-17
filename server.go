package main

import (
	"belajar-echo/config"
	"belajar-echo/router"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	config.Migrate()

	e := echo.New()

	router.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))

}
