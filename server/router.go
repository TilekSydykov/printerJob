package server

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/pagecount", PagecountHandler)
	r.HandleFunc("/lowtoner", LowtonerHandler)
	r.HandleFunc("/command", CommandHandler)

	r.HandleFunc("/printer/search_local", SearchLocalHandler)
	
	return r
}
