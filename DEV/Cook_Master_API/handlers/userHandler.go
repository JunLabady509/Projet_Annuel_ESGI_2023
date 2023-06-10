package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gastroguru/cache"
	"gastroguru/user"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/asdine/storm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func bodyToUser(r *http.Request, u *user.User) error {
	if r == nil {
		return errors.New("a request is required")
	}

	if r.Body == nil {
		return errors.New("no request body")
	}

	if u == nil {
		return errors.New("an user is required")
	}

	bd, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)
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
		case sql.ErrNoRows:
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
		case sql.ErrNoRows:
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
