package handlers

import (
	"encoding/json"
	"errors"
	"gastroguru/cache"
	"gastroguru/user"
	"io"
	"net/http"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
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

	u.ID = bson.NewObjectId()
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
	w.Header().Set("Location", "/users/"+u.ID.Hex())
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

func usersGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {

	if cache.Serve(w, r) {
		return
	}

	u, err := user.One(id)
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

func usersPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u := new(user.User)
	err := bodyToUser(r, u)
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

func usersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	u, err := user.One(id)
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

func usersDeleteOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	err := user.Delete(id)
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
