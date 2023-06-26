package handlers

import (
	"database/sql"
	"fmt"
	"gastroguru/cache"
	"gastroguru/database"
	"gastroguru/learning"

	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func workshopsPostOne(ctx echo.Context) error {
	c := new(learning.Workshop)
	err := ctx.Bind(c)
	if err != nil {
		fmt.Println("Error:", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = c.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			fmt.Println("Error saving")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cache.Drop("/workshops")
	ctx.Response().Header().Set("Location", "/workshops/"+strconv.Itoa(c.ID))
	return ctx.NoContent(http.StatusCreated)
}

func workshopsGetAll(ctx echo.Context) error {
	workshops, err := learning.GetAllWorkshops()
	if err != nil {
		fmt.Println("Error getting all workshops")
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, jsonResponse{"workshops": workshops})
}

func workshopsGetOne(ctx echo.Context) error {
	if cache.Serve(ctx.Response(), ctx.Request()) {
		return nil
	}

	id := ctx.Param("id")

	c, err := learning.GetWorkshop(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.JSON(http.StatusOK, jsonResponse{"workshop": c})
}

func workshopsPutOne(ctx echo.Context) error {
	c := new(learning.Workshop)
	err := ctx.Bind(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = c.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/workshops")
	return ctx.JSON(http.StatusOK, jsonResponse{"workshop": c})

}

func workshopsPatchOne(ctx echo.Context) error {

	id := ctx.Param("id")

	c, err := learning.GetWorkshop(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	err = ctx.Bind(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	c.ID, _ = strconv.Atoi(id)

	err = c.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/workshops")
	return ctx.JSON(http.StatusOK, jsonResponse{"workshop": c})

}

func workshopsDeleteOne(ctx echo.Context) error {

	id := ctx.Param("id")

	err := learning.DeleteWorkshop(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	cache.Drop("/workshops")
	cache.Drop(cache.MakeResource(ctx.Request()))
	return ctx.NoContent(http.StatusOK)

}

func workshopsOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func workshopOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}
