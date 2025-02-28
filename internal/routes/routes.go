package routes

import (
	"net/http"

	"github.com/erik1502/Cloud/web/templates/pages"
	"github.com/erik1502/Cloud/web/templates/util"
	"github.com/labstack/echo/v4"
)

type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type Register struct {
	Username        string `form:"username"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

// username and password
var db = make(map[string]string)

func RegisterRoutes(e *echo.Echo) {
	protected := e.Group("", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if _, err := c.Cookie("username"); err != nil {
				return c.Redirect(http.StatusFound, "/login")
			}

			return next(c)
		}
	})

	protected.Use()

	protected.GET("/", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.HomePage())
	})

	e.GET("/login", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.LoginPage())
	})

	e.POST("/login", func(c echo.Context) error {
		var l Login
		if err := c.Bind(&l); err != nil {
			return c.String(http.StatusInternalServerError, "Something bad happend on our end :(")
		}

		password, ok := db[l.Username]
		if !ok || password != l.Password {
			return c.String(http.StatusUnauthorized, "Wrong username or password.")
		}

		// Session
		mockSession(c, l.Username)

		return c.Redirect(http.StatusSeeOther, "/")
	})

	e.GET("/import", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.ImportForm())
	})

	e.GET("/register", func(c echo.Context) error {
		return util.Render(c, http.StatusOK, pages.RegisterPage())
	})

	e.POST("/register", func(c echo.Context) error {
		var r Register
		if err := c.Bind(&r); err != nil {
			return c.String(http.StatusInternalServerError, "Something bad happend on our end :(")
		}

		_, ok := db[r.Username]
		if ok {
			return c.String(http.StatusConflict, "User name is not available.")
		}

		if r.Password != r.ConfirmPassword {
			return c.String(http.StatusUnprocessableEntity, "Passwords do not match.")
		}

		// Save User credentials
		db[r.Username] = r.Password

		// Session
		mockSession(c, r.Username)

		return c.Redirect(http.StatusSeeOther, "/")
	})

	e.Any("*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "404 - Not Found")
	})
}

func mockSession(c echo.Context, username string) {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = username
	cookie.Secure = true

	c.SetCookie(cookie)
}
