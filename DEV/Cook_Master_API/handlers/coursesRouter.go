package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitCoursesRouter(e *echo.Echo) {
	c := e.Group("/courses")

	c.OPTIONS("", coursesOptions)
	c.HEAD("", coursesGetAll, cache.ServeCache)
	c.GET("", coursesGetAll, cache.ServeCache, cache.CacheResponse)
	c.POST("", coursesPostOne) //middleware.BasicAuth(auth))

	cid := c.Group("/:id")

	cid.GET("", coursesGetOne, cache.ServeCache, cache.CacheResponse)
	cid.OPTIONS("", courseOptions)
	cid.HEAD("", coursesGetOne, cache.ServeCache)
	cid.PUT("", coursesPutOne, middleware.BasicAuth(auth), cache.CacheResponse)
	cid.PATCH("", coursesPatchOne, middleware.BasicAuth(auth), cache.CacheResponse)
	cid.DELETE("", coursesDeleteOne, middleware.BasicAuth(auth))
}
