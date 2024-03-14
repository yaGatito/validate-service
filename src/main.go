package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"validate-service/src/handlers"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	http.HandleFunc("/validate", handlers.HandleValidatePath)

	port := os.Getenv("VALIDATE_SERVICE_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is started on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
