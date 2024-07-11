package main

import (
	"net/http"
	"os"

	"github.com/gocroot/route"
)

func main() {
	http.HandleFunc("/", route.URL)
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
