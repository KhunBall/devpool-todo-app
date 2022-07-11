package main

import (
	db "goapi/db"
	u "goapi/users"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.DB()
	db.Migrate()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"http://localhost:4200"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"response": "success"}`)
	})

	e.GET("/users", u.GetAll)
	e.GET("/users/email/:email", u.GetbyMail)
	e.GET("/users/:id", u.GetUser)
	e.POST("/users", u.Save)
	// e.PUT("/users/:id", u.Update)
	e.DELETE("/users/:id", u.Delete)

	// e.Logger.Fatal(e.Start(":80"))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))

}
