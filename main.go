package main

import (
	// "log"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
	"github.com/erik1502/Cloud/web/templates/pages"
	"github.com/erik1502/Cloud/web/templates/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

	e.GET("/", func(c echo.Context) error {
        return util.Render(c, http.StatusOK, pages.HomePage())
	})

	// Serve static files
	e.Static("/static", "web/static")

	// Start the Server
	e.Logger.Fatal(e.Start(":" + PORT))
}
