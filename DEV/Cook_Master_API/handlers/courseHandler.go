package handlers

import (
	"database/sql"
	"fmt"
	"gastroguru/cache"
	course "gastroguru/courses"
	"gastroguru/database"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// func bodyToCourse(r *http.Request, c *course.Course) error {
// 	if r == nil {
// 		return errors.New("a request is required")
// 	}

// 	if r.Body == nil {
// 		return errors.New("no request body")
// 	}

// 	if c == nil {
// 		return errors.New("an user is required")
// 	}

// 	bd, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal(bd, c)
// }

func coursesPostOne(ctx echo.Context) error {
	c := new(course.Course)
	err := ctx.Bind(c)
	if err != nil {
		fmt.Println("Error Binding")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = c.Save()

	if err != nil {
		if err == database.ErrRecordInvalid {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ctx.Response().Header().Set("Location", "/courses/"+strconv.Itoa(c.ID))
	return ctx.NoContent(http.StatusCreated)
}

func coursesGetAll(ctx echo.Context) error {
	courses, err := course.GetAllCourses()
	if err != nil {
		fmt.Println("Error getting all courses")
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, jsonResponse{"courses": courses})
}

func coursesGetOne(ctx echo.Context) error {
	if cache.Serve(ctx.Response(), ctx.Request()) {
		return nil
	}

	id := ctx.Param("id")

	c, err := course.GetCourse(id)
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
	return ctx.JSON(http.StatusOK, jsonResponse{"course": c})
}

func coursesPutOne(ctx echo.Context) error {
	c := new(course.Course)
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
	cache.Drop("/courses")
	return ctx.JSON(http.StatusOK, jsonResponse{"course": c})

}

func coursesPatchOne(ctx echo.Context) error {

	id := ctx.Param("id")

	c, err := course.GetCourse(id)
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
	cache.Drop("/users")
	return ctx.JSON(http.StatusOK, jsonResponse{"course": c})

}

func coursesDeleteOne(ctx echo.Context) error {

	id := ctx.Param("id")

	err := course.DeleteCourse(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	cache.Drop("/courses")
	cache.Drop(cache.MakeResource(ctx.Request()))
	return ctx.NoContent(http.StatusOK)

}

func coursesOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func courseOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}
