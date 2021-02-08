package server

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func SearchLocalHandler(w http.ResponseWriter, r *http.Request){
	dat,_:= ioutil.ReadFile("printer.configs")
	fmt.Fprintf(w, string(dat))
}

