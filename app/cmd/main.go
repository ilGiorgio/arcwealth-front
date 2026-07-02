package main

import (
	"fmt"
	"os"

	"app/internal/api"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (c *CustomValidator) Validate(i any) error {
	if err := c.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Static("/images", "images")
	e.Static("/css", "css")
	// Serve images from outside the project working from docker
	e.Static("/assets", "assets")

	// Register API Routes from separate package
	api.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":42069"))
}
