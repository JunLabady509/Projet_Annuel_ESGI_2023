package main

import (
	"fmt"
	"gastroguru/handlers"
	"net/http"
	"os"
)

func main() {

	// user.TestUserFunctions(nil)

	// handlers.TestPostBodyResponse(nil)

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	fmt.Println("Server started on port 44444")
	err := http.ListenAndServe("localhost:44444", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
