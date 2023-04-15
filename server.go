package main

import (
	"bpjs-mono/helpers"
	"bpjs-mono/models"
	"bpjs-mono/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errLoadingEnvFile := godotenv.Load()
	if errLoadingEnvFile != nil {
		helpers.HandleError("error loading the .env file", errLoadingEnvFile)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	models.ModelsDb = models.Connect(1)
	routes.Build(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

}