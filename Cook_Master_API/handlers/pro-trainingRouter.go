package handlers

import (
	"gastroguru/cache"

	"github.com/labstack/echo"
)

func InitProfessionalTrainingsRouter(e *echo.Echo) {
	c := e.Group("/professional_trainings")

	c.OPTIONS("", professionalTrainingsOptions)
	c.HEAD("", professionalTrainingsGetAll, cache.ServeCache)
	c.GET("", professionalTrainingsGetAll, cache.ServeCache, cache.CacheResponse)
	c.POST("", professionalTrainingsPostOne) //middleware.BasicAuth(auth))

	cid := c.Group("/:id")

	cid.GET("", professionalTrainingsGetOne, cache.ServeCache, cache.CacheResponse)
	cid.OPTIONS("", professionalTrainingOptions)
	cid.HEAD("", professionalTrainingsGetOne, cache.ServeCache)
	cid.PUT("", professionalTrainingsPutOne /*middleware.BasicAuth(auth),*/, cache.CacheResponse)
	cid.PATCH("", professionalTrainingsPatchOne /*middleware.BasicAuth(auth),*/, cache.CacheResponse)
	cid.DELETE("", professionalTrainingsDeleteOne /*middleware.BasicAuth(auth),*/)
}
