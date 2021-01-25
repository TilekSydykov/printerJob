package server

import (
	"fmt"
	"net/http"
	"printsServer/util"
	"strings"
	"strconv"
	"encoding/json"
)

type PageCountMessage struct {
	PrintedPages int
}

type ErrorMessage struct {
	Error string
}

type StatusMessage struct {
	Code string
	Display string
	Online string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StatusHandler(w http.ResponseWriter, r *http.Request){
	p := util.GetStatus()
	stringSlice := strings.Split(p, "\n")
	sm := StatusMessage{"", "", ""}
	for _, s := range stringSlice{
		m, n := parseString(s)
		if m == "CODE"{
			sm.Code = n
		}
		if m == "DISPLAY" {
			sm.Display = n
		}
		if m == "ONLINE" {
			sm.Online = n
		}
	}
	res, _ := json.Marshal(sm)
	fmt.Fprintf(w, string(res))
}

func PagecountHandler(w http.ResponseWriter, r *http.Request){
	p := util.GetPageCount()

	p = strings.ReplaceAll(p, " ", "")
	p = strings.ReplaceAll(p, "\n", "")
	p = strings.ReplaceAll(p, "\r", "")
	p = strings.ReplaceAll(p, "\f", "")
	h, err := strconv.Atoi(p)
	
	if err != nil{
		res, _ := json.Marshal(ErrorMessage{"converting problem. value = " + p})
		fmt.Fprintf(w, string(res))
		return
	}
	res, _ := json.Marshal(PageCountMessage{h})
	fmt.Fprintf(w, string(res))
}
