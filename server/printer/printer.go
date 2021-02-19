package printer

import (
	"encoding/json"
	"fmt"
	"github.com/signintech/gopdf"
	"net/http"
	"github.com/TilekSydykov/printsServer/util"
)

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
