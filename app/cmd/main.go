package main

import (
	"app/internal/api"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/images", "images" )
	e.Static("/css", "css" )
	// Serve images from outside the project working from docker
	e.Static("/assets", "assets")

	// Register API Routes from separate package
	api.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":42069"))
}
