package main

import (
	"fmt"
	"log"
	"net/http"
	"printsServer/server"
)

func main() {
	fmt.Println("Server working on port 8080")
	http.Handle("/", server.GetRouter())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
