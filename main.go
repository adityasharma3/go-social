package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adityasharma3/go-social/router"
)

func main() {
	fmt.Println("Hello - this app works now")
	router := router.Router()

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
