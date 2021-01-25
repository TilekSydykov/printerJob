package server

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/pagecount", PagecountHandler)
	return r
}
