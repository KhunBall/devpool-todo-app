package main

import (
	db "goapi/db"
	u "goapi/users"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.DB()
	db.Migrate()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().Host, "192.168.174.172") {
				return true
			}
			return false
		},
	}))
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// AllowOrigins: []string{"http://localhost"},
	// AllowOrigins: []string{"*"},
	// AllowOrigins: []string{"*", "http://192.168.174.172:4200"},
	// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization },
	// AllowMethods: []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"response": "success"}`)
	})

	e.GET("/users", u.GetAll)
	e.GET("/users/email/:email", u.GetbyMail)
	e.GET("/users/:id", u.GetUser)
	e.POST("/users", u.Save)
	// e.PUT("/users/:id", u.Update)
	e.DELETE("/users/:id", u.Delete)

	e.Logger.Fatal(e.Start(":1234"))

	// httpPort := os.Getenv("HTTP_PORT")
	// if httpPort == "" {
	// 	httpPort = "1234"
	// }

	// e.Logger.Fatal(e.Start(":" + httpPort))

}
