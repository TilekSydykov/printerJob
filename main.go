package main

import (
	"fmt"
	"log"
	"net/http"
	"printsServer/server"
	"printsServer/util"
)

func main() {
	go beginServer()
	fmt.Println("Server working on port 8080")
	http.Handle("/", server.GetRouter())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func beginServer() {
	fmt.Println(util.GetPageCount())
}
