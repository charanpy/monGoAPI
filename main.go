package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/charanpy/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Server started")
}
