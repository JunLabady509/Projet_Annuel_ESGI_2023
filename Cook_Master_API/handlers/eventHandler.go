package handlers

import (
	"database/sql"
	"fmt"
	"gastroguru/cache"
	"gastroguru/database"
	"gastroguru/events"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func eventsPostOne(ctx echo.Context) error {
	e := new(events.Event)
	err := ctx.Bind(e)
	if err != nil {
		fmt.Println("Error:", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = e.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			fmt.Println("Error saving")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cache.Drop("/events")
	ctx.Response().Header().Set("Location", "/events/"+strconv.Itoa(e.ID))
	return ctx.NoContent(http.StatusCreated)
}

func eventsGetAll(ctx echo.Context) error {
	events, err := events.GetAllEvents()
	if err != nil {
		fmt.Println("Error getting all events")
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, jsonResponse{"events": events})
}

func eventsGetOne(ctx echo.Context) error {
	if cache.Serve(ctx.Response(), ctx.Request()) {
		return nil
	}

	id := ctx.Param("id")

	e, err := events.GetEvent(id)
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
	return ctx.JSON(http.StatusOK, jsonResponse{"event": e})
}

func eventsPutOne(ctx echo.Context) error {
	e := new(events.Event)
	err := ctx.Bind(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = e.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/events")
	return ctx.JSON(http.StatusOK, jsonResponse{"event": e})

}

func eventsPatchOne(ctx echo.Context) error {

	id := ctx.Param("id")

	e, err := events.GetEvent(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	err = ctx.Bind(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	e.ID, _ = strconv.Atoi(id)

	err = e.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/events")
	return ctx.JSON(http.StatusOK, jsonResponse{"event": e})

}

func eventsDeleteOne(ctx echo.Context) error {

	id := ctx.Param("id")

	err := events.DeleteEvent(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	cache.Drop("/events")
	cache.Drop(cache.MakeResource(ctx.Request()))
	return ctx.NoContent(http.StatusOK)

}

func eventsOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func eventOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}
