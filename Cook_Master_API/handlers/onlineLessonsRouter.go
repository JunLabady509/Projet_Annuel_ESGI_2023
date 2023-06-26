package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
)

func InitOnlineCoursesRouter(e *echo.Echo) {

	oc := e.Group("/onlinelessons")

	oc.GET("", onlineLessonsGetAll, cache.ServeCache, cache.CacheResponse)
	oc.POST("", onlineLessonsPostOne)
	oc.HEAD("", onlineLessonsGetAll, cache.ServeCache)
	oc.OPTIONS("", onlineLessonsOptions)

	ocid := oc.Group("/:id")

	ocid.GET("", onlineLessonsGetOne, cache.ServeCache, cache.CacheResponse)
	ocid.HEAD("", onlineLessonsGetOne, cache.ServeCache)
	ocid.OPTIONS("", onlineLessonOptions)
	ocid.PUT("", onlineLessonsPutOne, cache.CacheResponse)
	ocid.PATCH("", onlineLessonsPatchOne, cache.CacheResponse)
	ocid.DELETE("", onlineLessonsDeleteOne)
}
