package handlers

import (
	"database/sql"
	"fmt"
	"gastroguru/cache"
	"gastroguru/database"
	"gastroguru/homecourses"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func homeCoursePostOne(ctx echo.Context) error {
	h := new(homecourses.HomeCourse)
	err := ctx.Bind(h)
	if err != nil {
		fmt.Println("Error Binding :", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = h.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cache.Drop("/homecourses")
	ctx.Response().Header().Set("Location", "/homecourses/"+strconv.Itoa(h.ID))
	return ctx.NoContent(http.StatusCreated)
}

func homeCourseGetAll(ctx echo.Context) error {
	homeCourses, err := homecourses.GetAllHomeCourses()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.JSON(http.StatusOK, jsonResponse{"homecourses": homeCourses})
}

func homeCourseGetOne(ctx echo.Context) error {
	if cache.Serve(ctx.Response(), ctx.Request()) {
		return nil
	}

	id := ctx.Param("id")

	h, err := homecourses.GetHomeCourse(id)

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
	return ctx.JSON(http.StatusOK, jsonResponse{"homecourse": h})
}

func homeCoursePutOne(ctx echo.Context) error {
	h := new(homecourses.HomeCourse)
	err := ctx.Bind(h)
	if err != nil {
		fmt.Println("Error Binding :", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = h.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/homecourses")
	return ctx.NoContent(http.StatusCreated)
}

func homeCoursePatchOne(ctx echo.Context) error {

	id := ctx.Param("id")

	h, err := homecourses.GetHomeCourse(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	err = ctx.Bind(h)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	h.ID, _ = strconv.Atoi(id)

	err = h.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cache.Drop("/homecourses")
	return ctx.JSON(http.StatusOK, jsonResponse{"homecourses": h})

}

func homeCourseDeleteOne(ctx echo.Context) error {
	id := ctx.Param("id")

	err := homecourses.DeleteHomeCourse(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	cache.Drop("/homecourses")
	cache.Drop(cache.MakeResource(ctx.Request()))
	return ctx.NoContent(http.StatusOK)
}

func homeCoursesOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func homeCourseOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}
