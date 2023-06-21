package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "form.html") // Serve the HTML form
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Missing query", http.StatusBadRequest)
			return
		}

		// Run the CGI script
		cmd := exec.Command("./test", query)
		cmd.Stdout = w
		err := cmd.Run()
		if err != nil {
			log.Printf("Failed to run script: %v", err)
			http.Error(w, "Failed to run script", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
