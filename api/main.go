package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ChangeEnvironment()
	fmt.Println(config.Port)
	fmt.Println("Running api")

	r := router.Generate()

	fmt.Println(fmt.Sprintf("Server running on port %d", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
