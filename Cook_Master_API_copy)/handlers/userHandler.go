package handlers

import (
	"encoding/json"
	"errors"
	"gastroguru/cache"
	"gastroguru/user"
	"io"
	"net/http"
	"strconv"

	"github.com/asdine/storm"
	_ "github.com/go-sql-driver/mysql"
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

func usersPostOne(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	cache.Drop("/users")
	w.Header().Set("Location", "/users/"+strconv.Itoa(u.ID))
	w.WriteHeader(http.StatusCreated)

}

func usersGetAll(w http.ResponseWriter, r *http.Request) {

	if cache.Serve(w, r) {
		return
	}

	users, err := user.All()

	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == "HEAD" {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	cacheWriter := cache.NewWriter(w, r)
	postBodyResponse(cacheWriter, http.StatusOK, jsonResponse{"users": users})

}

func usersGetOne(w http.ResponseWriter, r *http.Request, id int) {

	if cache.Serve(w, r) {
		return
	}

	u, err := user.One(strconv.Itoa(id))
	if err != nil {
		switch err {
		case storm.ErrNotFound:
			postError(w, http.StatusNotFound)
			return
		default:
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	if r.Method == "HEAD" {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	cacheWriter := cache.NewWriter(w, r)
	postBodyResponse(cacheWriter, http.StatusOK, jsonResponse{"user": u})
}

func usersPutOne(w http.ResponseWriter, r *http.Request, id int) {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	cacheWriter := cache.NewWriter(w, r)
	postBodyResponse(cacheWriter, http.StatusOK, jsonResponse{"user": u})

}

func usersPatchOne(w http.ResponseWriter, r *http.Request, id int) {
	u, err := user.One(strconv.Itoa(id))
	if err != nil {
		switch err {
		case storm.ErrNotFound:
			postError(w, http.StatusNotFound)
			return
		default:
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	err = bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	u.ID = id
	err = u.Save()

	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	cache.Drop("/users")
	cacheWriter := cache.NewWriter(w, r)
	postBodyResponse(cacheWriter, http.StatusOK, jsonResponse{"user": u})

}

func usersDeleteOne(w http.ResponseWriter, r *http.Request, id int) {
	err := user.Delete(strconv.Itoa(int(id)))
	if err != nil {
		switch err {
		case storm.ErrNotFound:
			postError(w, http.StatusNotFound)
			return
		default:
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	cache.Drop("/users")
	cache.Drop(cache.MakeResource(r))
	w.WriteHeader(http.StatusOK)

}
