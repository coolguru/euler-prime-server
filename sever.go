package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/", AboutHandler)
	e.Get("/eulerprime", CalculateEulerPrimeHandler)
	e.POST("/eulerprimeupload", UploadAndCalculateEulerPrimeHandler)

	log.Println("Euler Prime server running at localhost:3600")
	e.Run(standard.New(":3600"))
}
