package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"github.com/TilekSydykov/printsServer/util"
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
	_, _ = fmt.Fprintf(w, "i am online")
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
		m, n := util.ParseString(s)
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

func CommandHandler(w http.ResponseWriter, r *http.Request) {
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

func GetMacHandler(w http.ResponseWriter, r *http.Request) {
	res, err := getMacAddr()
	if err != nil {
		_, _ = fmt.Fprintf(w, string("err"))
	}
	res = strings.Replace(res, ":", "", 10)
	res = strings.ToTitle(res)
	_, _ = fmt.Fprintf(w, res)
}

func getMacAddr() (string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	if len(ifas) > 0 {
		return ifas[0].HardwareAddr.String(), nil
	}
	return "", err
}
