package main

import (
	"log"
	"net/http"
	"validate-service/src/handlers"
)

func main() {
	http.HandleFunc("/validate", handlers.ValidateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//str1 := fmt.Sprintf("%03d", 100)
	//fmt.Println("String using fmt.Sprintf:", str1) // Output: 001
}
