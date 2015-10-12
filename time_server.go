package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Registrering the service
	http.HandleFunc(
		// Setting the url
		"/time",
		// Anonymous function
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
		})

	http.ListenAndServe(":8080", nil)
}
