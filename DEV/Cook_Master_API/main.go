package main

import (
	"gastroguru/database"
	"gastroguru/handlers"

	// database "gastroguru/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database.InitDB()
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${latency_human} \n",
	}))

	handlers.InitUsersRouter(e)
	handlers.InitCoursesRouter(e)

	e.Logger.Fatal(e.Start(":44446"))

}
