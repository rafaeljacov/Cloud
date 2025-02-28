package main

import (
	// "log"

	"os"

	// "github.com/joho/godotenv"

	"github.com/erik1502/Cloud/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MockUserStore struct {
	username string
	password string
}

func main() {
	// Load environment variables from .env file
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	// Read the PORT from the environment
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default to 8080 if not set
	}

    routes.RegisterRoutes(e)

	// Serve static files
	e.Static("/static", "web/static")

	// Start the Server
	e.Logger.Fatal(e.Start(":" + PORT))
}
