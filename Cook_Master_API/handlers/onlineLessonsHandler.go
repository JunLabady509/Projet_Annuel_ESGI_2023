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

func onlineLessonsPostOne(ctx echo.Context) error {

	ol := new(learning.OnlineLesson)
	err := ctx.Bind(ol)
	if err != nil {
		fmt.Println("Error Binding")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = ol.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			fmt.Println("Error saving")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ctx.Response().Header().Set("Location", "/workshops/"+strconv.Itoa(ol.ID))
	return ctx.NoContent(http.StatusCreated)
}

func onlineLessonsGetAll(ctx echo.Context) error {
	onlineLessons, err := learning.GetAllOnlineLessons()
	if err != nil {
		fmt.Println("Error getting all Online Courses")
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, jsonResponse{"online_lessons": onlineLessons})
}

func onlineLessonsGetOne(ctx echo.Context) error {
	if cache.Serve(ctx.Response(), ctx.Request()) {
		return nil
	}

	id := ctx.Param("id")

	ol, err := learning.GetOnlineLesson(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, jsonResponse{"online_lesson": ol})
}

func onlineLessonsPutOne(ctx echo.Context) error {
	ol := new(learning.OnlineLesson)
	err := ctx.Bind(ol)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = ol.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/onlinelessons/" + strconv.Itoa(ol.ID))
	return ctx.JSON(http.StatusOK, jsonResponse{"online_lesson": ol})

}

func onlineLessonsPatchOne(ctx echo.Context) error {

	id := ctx.Param("id")

	ol, err := learning.GetOnlineLesson(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	err = ctx.Bind(ol)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ol.ID, _ = strconv.Atoi(id)

	err = ol.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/users")
	return ctx.JSON(http.StatusOK, jsonResponse{"online_lesson": ol})

}

func onlineLessonsDeleteOne(ctx echo.Context) error {
	id := ctx.Param("id")

	err := learning.DeleteOnlineLesson(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	cache.Drop("/onlinelessons")
	cache.Drop(cache.MakeResource(ctx.Request()))
	return ctx.NoContent(http.StatusNoContent)
}

func onlineLessonsOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func onlineLessonOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}
