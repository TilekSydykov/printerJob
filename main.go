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

	fmt.Print("Server working on port 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", server.MainHandler)
	r.HandleFunc("/status", server.StatusHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func beginServer(){
	_ =  util.GetConn()

}


