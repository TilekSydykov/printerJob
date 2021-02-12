package printer

import (
	"fmt"
	"github.com/signintech/gopdf"
	"net/http"
)

func PrintImage(w http.ResponseWriter, r *http.Request) {
	imageId, err := r.URL.Query()["img_id"]
	if !err || len(imageId) == 0 {
		fmt.Fprintf(w, "no id")
		return
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	pdf.Image("/home/terminal/scanned_images/images"+imageId[0]+".jpg", 200, 50, nil)
	pdf.WritePdf("/home/terminal/scanned_images/images/image.pdf")
}
