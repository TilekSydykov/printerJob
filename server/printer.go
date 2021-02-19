package server

import (
	"encoding/json"
	"fmt"
	"github.com/signintech/gopdf"
	"net"
	"net/http"
	"printsServer/config"
	"printsServer/util"
	"strconv"
	"time"
)

func SearchLocalHandler(w http.ResponseWriter, r *http.Request) {
	var conf config.Config
	var req = struct {
		Gate string `json:"gate"`
	}{""}
	err := conf.GetConfig()
	if err != nil {
		_ = json.NewEncoder(w).Encode(ErrorMessage{err.Error()})
	}
	err = json.NewDecoder(r.Body).Decode(req)
	if req.Gate != "" {
		conf.LocalGate = req.Gate
	}

	for i := 0; i < 256; i++ {
		go callToPrinter(conf, i)
	}
}

func callToPrinter(conf config.Config, num int) {
	_, err := net.DialTimeout("tcp", conf.LocalGate+string(num)+":"+config.PrinterPort, 1*time.Second)
	if err != nil {
		return
	}
	conf.Ip = conf.LocalGate + string(num)
	err = conf.WriteConfig()
	fmt.Print(conf.LocalGate + strconv.Itoa(num) + "error\n")
	if err != nil {
		print("cannot write")
	}
}

func PrintImage(w http.ResponseWriter, r *http.Request) {
	imageId, err := r.URL.Query()["img_id"]
	if !err || len(imageId) == 0 {
		fmt.Fprintf(w, "no id")
		return
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeB4})
	pdf.AddPage()
	pdf.Image("/home/terminal/scanned_images/images/"+imageId[0]+".jpg", 0, 0, gopdf.PageSizeA4)
	e := pdf.WritePdf("/home/terminal/scanned_images/images/image.pdf")
	if e != nil {
		fmt.Fprintf(w, e.Error())
		return
	}
	_, _ = fmt.Fprintf(w, "image.pdf")
}

type PrintRequest struct {
	Path string
}

func PrintPdf(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request PrintRequest
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	err = util.PrintDoc(request.Path, 0, false)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "printing")
}
