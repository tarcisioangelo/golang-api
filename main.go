package main

import (
	"api/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.Init()

	fmt.Println("Server booting on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
