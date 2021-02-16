package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/pagecount", PagecountHandler)
	r.HandleFunc("/lowtoner", LowtonerHandler)
	r.HandleFunc("/command", CommandHandler)

	r.HandleFunc("/getmac", GetMacHandler)

	s := r.PathPrefix("/printer").Subrouter()
	s.HandleFunc("/search_local", SearchLocalHandler)
	s.HandleFunc("/image", PrintImage)

	s.HandleFunc("/print", PrintPdf)

	fs := http.FileServer(http.Dir("/home/terminal/scanned_images"))
	r.PathPrefix("/images/").Handler(fs)

	return r
}
