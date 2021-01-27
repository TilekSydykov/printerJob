package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"printsServer/util"
	"strconv"
	"strings"
)

type PageCountMessage struct {
	PrintedPages int
}

type CommandResponce struct {
	Responce string
}

type ErrorMessage struct {
	Error string
}

type StatusMessage struct {
	Code    string
	Display string
	Online  string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	p, err := util.GetStatus()
	if err != nil {
		res, _ := json.Marshal(ErrorMessage{err.Error()})
		_, _ = fmt.Fprintf(w, string(res))
		return
	}
	stringSlice := strings.Split(p, "\n")
	sm := StatusMessage{"", "", ""}
	for _, s := range stringSlice {
		m, n := parseString(s)
		if m == "CODE" {
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
	_, _ = fmt.Fprintf(w, string(res))
}

func CommandHandler(w http.ResponseWriter, r *http.Request){
	p, err := util.RunCommand(r.PostForm.Get("command"))
	if err != nil {
		res, _ := json.Marshal(ErrorMessage{err.Error()})
		_, _ = fmt.Fprintf(w, string(res))
		return
	}
	res, _ := json.Marshal(CommandResponce{p})
	_, _ = fmt.Fprintf(w, string(res))
}

func PagecountHandler(w http.ResponseWriter, r *http.Request) {
	p, err := util.GetPageCount()
	if err != nil {
		res, _ := json.Marshal(ErrorMessage{err.Error()})
		_, _ = fmt.Fprintf(w, string(res))
		return
	}
	p = strings.ReplaceAll(p, " ", "")
	p = strings.ReplaceAll(p, "\n", "")
	p = strings.ReplaceAll(p, "\r", "")
	p = strings.ReplaceAll(p, "\f", "")
	h, err := strconv.Atoi(p)

	if err != nil {
		res, _ := json.Marshal(ErrorMessage{"converting problem. value = " + p})
		_, _ = fmt.Fprintf(w, string(res))
		return
	}
	res, _ := json.Marshal(PageCountMessage{h})
	_, _ = fmt.Fprintf(w, string(res))
}

func LowtonerHandler(w http.ResponseWriter, r *http.Request) {
	p, err := util.Gettoner()
	if err != nil {
		res, _ := json.Marshal(ErrorMessage{err.Error()})
		_, _ = fmt.Fprintf(w, string(res))
		return
	}
	print(p)
	h := ErrorMessage{p}
	res, _ := json.Marshal(h)
	_, _ = fmt.Fprintf(w, string(res))
}
