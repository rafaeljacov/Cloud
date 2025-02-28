package routes

import (
	"net/http"

	"github.com/erik1502/Cloud/web/templates/pages"
	"github.com/erik1502/Cloud/web/templates/util"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.HomePage())
	})

	e.GET("/login", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.LoginPage())
	})

	e.GET("/register", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.RegisterPage())
	})
}
