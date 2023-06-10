package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouter(e *echo.Echo) {

	e.GET("/", root)

	u := e.Group("/users")

	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll, cache.ServeCache)
	u.GET("", usersGetAll, cache.ServeCache, cache.CacheResponse)
	u.POST("", usersPostOne) //, middleware.BasicAuth(auth))

	uid := u.Group("/:id")

	uid.GET("", usersGetOne, cache.ServeCache, cache.CacheResponse)
	uid.OPTIONS("", userOptions)
	uid.HEAD("", usersGetOne, cache.ServeCache)
	uid.PUT("", usersPutOne, middleware.BasicAuth(auth), cache.CacheResponse)
	uid.PATCH("", usersPatchOne, middleware.BasicAuth(auth), cache.CacheResponse)
	uid.DELETE("", usersDeleteOne, middleware.BasicAuth(auth))

	e.Logger.Fatal(e.Start(":44446"))

}
