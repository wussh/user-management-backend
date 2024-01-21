package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/register", register)
	e.POST("/login", login)

	// Start server
	e.Start(":8080")
}

func register(c echo.Context) error {
	return c.String(http.StatusOK, "User registered successfully")
}

func login(c echo.Context) error {

	return c.String(http.StatusOK, "User logged in successfully")
}
