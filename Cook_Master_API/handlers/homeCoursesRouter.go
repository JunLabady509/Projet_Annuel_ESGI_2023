package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
)

/*
Il reste à faire:
Configurer le middleware BasicAuth pour certaines routes
Comme pour : PUT, PATCH, DELETE
Mettre en place le système d'Autorisation pour les routes
en fonction du rôle de l'utilisateur

*/

func InitHomeCoursesRouter(e *echo.Echo) {
	h := e.Group("/homecourses")

	h.GET("", homeCourseGetAll, cache.ServeCache, cache.CacheResponse)
	h.POST("", homeCoursePostOne)
	h.HEAD("", homeCourseGetAll, cache.ServeCache)
	h.OPTIONS("", homeCoursesOptions)

	hid := h.Group("/:id")

	hid.GET("", homeCourseGetOne, cache.ServeCache, cache.CacheResponse)
	hid.HEAD("", homeCourseGetOne, cache.ServeCache)
	hid.OPTIONS("", homeCourseOptions)
	hid.PUT("", homeCoursePutOne, cache.CacheResponse)
	hid.PATCH("", homeCoursePatchOne, cache.CacheResponse)
	hid.DELETE("", homeCourseDeleteOne)
}
