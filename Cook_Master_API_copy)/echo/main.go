package main

import (
	"database/sql"
	"fmt"
	"gastroguru/cache"
	"gastroguru/user"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type jsonResponse map[string]interface{}

func serveCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if cache.Serve(ctx.Response(), ctx.Request()) {
			return nil
		}
		return next(ctx)
	}
}

func cacheResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Writer = cache.NewWriter(ctx.Response().Writer, ctx.Request())
		return next(ctx)
	}
}

func auth(username, password string, ctx echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}

func usersGetAll(ctx echo.Context) error {
	users, err := user.All()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.JSON(http.StatusOK, jsonResponse{"users": users})

}

func usersPostOne(ctx echo.Context) error {
	u := new(user.User)
	err := ctx.Bind(u)
	if err != nil {
		fmt.Println("Error Binding :", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cache.Drop("/users")
	ctx.Response().Header().Set("Location", "/users/"+strconv.Itoa(u.ID))
	return ctx.NoContent(http.StatusCreated)
}

func usersGetOne(ctx echo.Context) error {

	if cache.Serve(ctx.Response(), ctx.Request()) {
		return nil
	}

	id := ctx.Param("id")

	u, err := user.One(id)
	if err != nil {
		switch err {
		case storm.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.JSON(http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(ctx echo.Context) error {
	u := new(user.User)
	err := ctx.Bind(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	return ctx.JSON(http.StatusOK, jsonResponse{"user": u})

}

func usersPatchOne(ctx echo.Context) error {

	id := ctx.Param("id")

	u, err := user.One(id)
	if err != nil {
		switch err {
		case storm.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	err = ctx.Bind(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	u.ID, _ = strconv.Atoi(id)

	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	return ctx.JSON(http.StatusOK, jsonResponse{"user": u})

}

func usersDeleteOne(ctx echo.Context) error {

	id := ctx.Param("id")

	err := user.Delete(id)
	if err != nil {
		switch err {
		case storm.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	cache.Drop("/users")
	cache.Drop(cache.MakeResource(ctx.Request()))
	return ctx.NoContent(http.StatusOK)

}

func usersOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func userOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Welcome to Gastroguru")
}

func main() {

	var db *sql.DB
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${latency_human} \n",
	}))

	e.GET("/", root)

	u := e.Group("/users")

	u.OPTIONS("", usersOptions)
	u.HEAD("", usersGetAll, serveCache)
	u.GET("", usersGetAll, serveCache, cacheResponse)
	u.POST("", usersPostOne) //, middleware.BasicAuth(auth))

	uid := u.Group("/:id")

	uid.GET("", usersGetOne, serveCache, cacheResponse)
	uid.OPTIONS("", userOptions)
	uid.HEAD("", usersGetOne, serveCache)
	uid.PUT("", usersPutOne, middleware.BasicAuth(auth), cacheResponse)
	uid.PATCH("", usersPatchOne, middleware.BasicAuth(auth), cacheResponse)
	uid.DELETE("", usersDeleteOne, middleware.BasicAuth(auth))

	e.Logger.Fatal(e.Start(":44446"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		db.Close()
		os.Exit(1)
	}()

}
