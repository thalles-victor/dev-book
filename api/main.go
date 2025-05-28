package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running api")

	r := router.Generate()

	fmt.Println("Serve running at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
