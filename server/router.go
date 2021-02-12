package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"printsServer/server/printer"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", MainHandler)
	r.HandleFunc("/status", StatusHandler)
	r.HandleFunc("/pagecount", PagecountHandler)
	r.HandleFunc("/lowtoner", LowtonerHandler)
	r.HandleFunc("/command", CommandHandler)

	r.HandleFunc("/printer/search_local", SearchLocalHandler)

	r.HandleFunc("/getmac", GetMacHandler)

	r.HandleFunc("/printer/image", printer.PrintImage)

	fs := http.FileServer(http.Dir("/home/terminal/scanned_images"))
	r.PathPrefix("/images/").Handler(fs)

	return r
}
