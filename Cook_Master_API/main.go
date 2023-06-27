package main

import (
	"gastroguru/database"
	"gastroguru/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database.InitDB()

	database.CreateAllTables(database.Db)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${latency_human} \n",
	}))

	handlers.InitUsersRouter(e)
	handlers.InitWorkshopsRouter(e)
	handlers.InitHomeCoursesRouter(e)
	handlers.InitOnlineCoursesRouter(e)
	handlers.InitProfessionalTrainingsRouter(e)
	handlers.InitEventsRouter(e)

	e.Logger.Fatal(e.Start(":44446"))

}
