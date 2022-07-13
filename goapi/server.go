package main

import (
	db "goapi/db"
	u "goapi/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.DB()
	db.Migrate()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ }))


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"response": "success"}`)
	})

	e.GET("/users", u.GetAll)
	e.GET("/users/email/:email", u.GetbyMail)
	e.GET("/users/:id", u.GetUser)
	e.POST("/users", u.Save)
	// dd
	// e.PUT("/users/:id", u.Update)
	e.DELETE("/users/:id", u.Delete)

	e.Logger.Fatal(e.Start(":1234"))

	// httpPort := os.Getenv("HTTP_PORT")
	// if httpPort == "" {
	// 	httpPort = "1234"
	// }

	// e.Logger.Fatal(e.Start(":" + httpPort))

}
