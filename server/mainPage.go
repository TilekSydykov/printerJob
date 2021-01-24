package server

import (
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StatusHandler(w http.ResponseWriter, r *http.Request){

}
