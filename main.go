package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kartikeyyadav7/mongoapi/routes"
)

func main() {
	fmt.Println("Hello from a proper api")
	r := routes.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000...")
}
