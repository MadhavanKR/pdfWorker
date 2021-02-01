package pdfservices

import (
	"github.com/jung-kurt/gofpdf"
	"strings"
)

func ConvertToPdf(pdfFileName string, images []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, imgPath := range images {
		pdf.AddPage()
		imgOptions := gofpdf.ImageOptions{
			AllowNegativePosition: true,
			ImageType:             strings.Split(imgPath, ".")[1],
			ReadDpi:               true,
		}
		pdf.RegisterImageOptions(imgPath, imgOptions)
		pdf.ImageOptions(imgPath, 0, 0, 200, 290, false, imgOptions, 0, "")
	}
	return pdf.OutputFileAndClose(pdfFileName)
}