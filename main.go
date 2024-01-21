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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User

func register(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	users = append(users, *u)

	return c.String(http.StatusOK, "User registered successfully")
}

func login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	for _, user := range users {
		if user.Username == u.Username && user.Password == u.Password {
			return c.String(http.StatusOK, "User logged in successfully")
		}
	}

	return c.String(http.StatusUnauthorized, "Invalid credentials")
}
