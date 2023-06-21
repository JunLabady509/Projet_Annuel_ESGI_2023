package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet: // GET
			//fmt.Println("GET")
			usersGetAll(w, r)
			return
		case http.MethodPost: // POST
			fmt.Println("POST")
			usersPostOne(w, r)
			return
		case http.MethodHead: // HEAD
			fmt.Println("HEAD")
			usersGetAll(w, r)
		case http.MethodOptions: // OPTIONS
			fmt.Println("OPTIONS")
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/users/")

	id, err := strconv.Atoi(path)
	if err != nil {
		postError(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet: // GET
		usersGetOne(w, r, id)
		return
	case http.MethodPut: // PUT
		usersPutOne(w, r, id)
		return
	case http.MethodPatch: // PATCH
		usersPatchOne(w, r, id)
		return
	case http.MethodDelete: // DELETE
		usersDeleteOne(w, r, id)
		return
	case http.MethodHead: // HEAD
		usersGetOne(w, r, id)
		return
	case http.MethodOptions: // OPTIONS
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodHead, http.MethodOptions}, nil)
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
