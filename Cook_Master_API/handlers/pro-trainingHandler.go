package handlers

import (
	"database/sql"
	"fmt"
	"gastroguru/database"
	"gastroguru/learning"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

func professionalTrainingsPostOne(ctx echo.Context) error {
	c := new(learning.ProfessionalTraining)
	err := ctx.Bind(c)
	if err != nil {
		fmt.Println("Error Binding")
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

	ctx.Response().Header().Set("Location", "/professional_trainings/"+strconv.Itoa(c.ID))
	return ctx.NoContent(http.StatusCreated)
}

func professionalTrainingsGetAll(ctx echo.Context) error {
	trainings, err := learning.GetAllProfessionalTrainings()
	if err != nil {
		fmt.Println("Error getting all professional trainings")
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if ctx.Request().Method == http.MethodHead {
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.JSON(http.StatusOK, jsonResponse{"professional_trainings": trainings})
}

func professionalTrainingsGetOne(ctx echo.Context) error {
	id := ctx.Param("id")

	c, err := learning.GetProfessionalTraining(id)
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
	return ctx.JSON(http.StatusOK, jsonResponse{"professional_training": c})
}

func professionalTrainingsPutOne(ctx echo.Context) error {
	c := new(learning.ProfessionalTraining)
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
	return ctx.JSON(http.StatusOK, jsonResponse{"professional_training": c})
}

func professionalTrainingsPatchOne(ctx echo.Context) error {
	id := ctx.Param("id")

	c, err := learning.GetProfessionalTraining(id)
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
	return ctx.JSON(http.StatusOK, jsonResponse{"professional_training": c})
}

func professionalTrainingsDeleteOne(ctx echo.Context) error {
	id := ctx.Param("id")

	err := learning.DeleteProfessionalTraining(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return echo.NewHTTPError(http.StatusNotFound)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return ctx.NoContent(http.StatusOK)
}

func professionalTrainingsOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}

func professionalTrainingOptions(ctx echo.Context) error {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}
	ctx.Response().Header().Set("Allow", strings.Join(methods, ","))
	return ctx.NoContent(http.StatusOK)
}
