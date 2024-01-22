package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *gorm.DB

func main() {
	// Initialize PostgreSQL connection using GORM
	dsn := "user=my_user dbname=my_database password=my_password host=postgres sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL:", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting DB:", err)
		return
	}
	defer sqlDB.Close()

	// Auto Migrate the User model
	if err := db.AutoMigrate(&User{}); err != nil {
		fmt.Println("Error auto migrating:", err)
		return
	}

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
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func register(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Create user using GORM
	err := db.Create(u).Error
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

	// Query the user using GORM
	var existingUser User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&existingUser).Error
	if err == nil {
		return c.String(http.StatusOK, "User logged in successfully")
	} else if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	} else {
		return err
	}
}
