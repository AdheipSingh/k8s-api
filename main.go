package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k8s-api/routes"
)

func main() {
	fmt.Println("HTTP server listening 8081")
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))

}
