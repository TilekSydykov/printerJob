package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"printsServer/server"
	"printsServer/util"
)

func main(){
	go beginServer()

	fmt.Println("Server working on port 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", server.MainHandler)
	r.HandleFunc("/status", server.StatusHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func beginServer(){
	
	fmt.Println(util.GetPageCount())
}


