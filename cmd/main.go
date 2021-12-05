package main

import (
	"fmt"
	"github.com/hulk0301/example-app-backend/internal/router"
	"log"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", router.HandleRequest)
	log.Printf("Server listening on http://localhost:%d/", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
