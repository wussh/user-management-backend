package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

func main() {
	// Initialize PostgreSQL connection
	connStr := "user=my_user dbname=my_database password=my_password host=postgres sslmode=disable"
	var err error
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL:", err)
		return
	}
	defer db.Close()

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

func register(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Insert user into the PostgreSQL database
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "User registered successfully")
}

func login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Query the PostgreSQL database for the user
	row := db.QueryRow("SELECT username, password FROM users WHERE username=$1 AND password=$2", u.Username, u.Password)

	var username, password string
	err := row.Scan(&username, &password)
	if err == nil {
		return c.String(http.StatusOK, "User logged in successfully")
	} else if err == sql.ErrNoRows {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	} else {
		return err
	}
}
