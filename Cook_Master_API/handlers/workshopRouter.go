package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
)

func InitWorkshopsRouter(e *echo.Echo) {
	c := e.Group("/workshops")

	c.OPTIONS("", workshopsOptions)
	c.HEAD("", workshopsGetAll, cache.ServeCache)
	c.GET("", workshopsGetAll, cache.ServeCache, cache.CacheResponse)
	c.POST("", workshopsPostOne) //middleware.BasicAuth(auth))

	cid := c.Group("/:id")

	cid.GET("", workshopsGetOne, cache.ServeCache, cache.CacheResponse)
	cid.OPTIONS("", workshopOptions)
	cid.HEAD("", workshopsGetOne, cache.ServeCache)
	cid.PUT("", workshopsPutOne /*middleware.BasicAuth(auth),*/, cache.CacheResponse)
	cid.PATCH("", workshopsPatchOne /*middleware.BasicAuth(auth),*/, cache.CacheResponse)
	cid.DELETE("", workshopsDeleteOne /*middleware.BasicAuth(auth),*/)
}
