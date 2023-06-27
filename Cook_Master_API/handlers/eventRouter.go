package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
)

func InitEventsRouter(e *echo.Echo) {
	c := e.Group("/events")

	c.OPTIONS("", eventsOptions)
	c.HEAD("", eventsGetAll, cache.ServeCache)
	c.GET("", eventsGetAll, cache.ServeCache, cache.CacheResponse)
	c.POST("", eventsPostOne)

	cid := c.Group("/:id")

	cid.GET("", eventsGetOne, cache.ServeCache, cache.CacheResponse)
	cid.OPTIONS("", eventOptions)
	cid.HEAD("", eventsGetOne, cache.ServeCache)
	cid.PUT("", eventsPutOne, cache.CacheResponse)
	cid.PATCH("", eventsPatchOne, cache.CacheResponse)
	cid.DELETE("", eventsDeleteOne)
}
