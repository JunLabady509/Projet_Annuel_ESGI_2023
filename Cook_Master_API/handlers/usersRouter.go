package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
)

func InitUsersRouter(e *echo.Echo) {

	e.GET("/", root)

	e.POST("/users/checkEmail", userCheckEmail)
	e.POST("/auth", userAuth)

	u := e.Group("/users")

	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll, cache.ServeCache)
	u.GET("", usersGetAll, cache.ServeCache, cache.CacheResponse)
	u.POST("", usersPostOne)

	uid := u.Group("/:id")

	uid.GET("", usersGetOne, cache.ServeCache, cache.CacheResponse)
	uid.OPTIONS("", userOptions)
	uid.HEAD("", usersGetOne, cache.ServeCache)
	uid.PUT("", usersPutOne /*middleware.BasicAuth(auth),*/, cache.CacheResponse)
	uid.PATCH("", usersPatchOne /*middleware.BasicAuth(auth),*/, cache.CacheResponse)
	uid.DELETE("", usersDeleteOne /*middleware.BasicAuth(auth)*/)
}
