package pdfservices

import (
	"errors"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
)

func ConvertToPdf(pdfFileName string, images []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, imgPath := range images {
		imgOptions := gofpdf.ImageOptions{
			AllowNegativePosition: true,
			ImageType:             strings.Split(imgPath, ".")[1],
			ReadDpi:               false,
		}
		h, w, imgDimensionErr := getImageDimensions(imgPath)
		if imgDimensionErr != nil {
			return imgDimensionErr
		}
		pageSize, pageSizeHeight, pageSizeWidth := getPageSize(convertPixelToMM(float64(h)), convertPixelToMM(float64(w)))
		if h > w {
			//adding potrait image
			log.Printf("adding a potrait page of type: %s\n", pageSize)
			pdf.AddPageFormat("P", gofpdf.SizeType{
				Ht: float64(pageSizeHeight),
				Wd: float64(pageSizeWidth),
			})
		} else {
			//adding landscape image
			log.Printf("adding a landscape page of type: %s\n", pageSize)
			pdf.AddPageFormat("L", gofpdf.SizeType{
				Ht: float64(pageSizeHeight),
				Wd: float64(pageSizeWidth),
			})
		}
		pdf.RegisterImageOptions(imgPath, imgOptions)
		pdf.ImageOptions(imgPath, 0, 0, -1, -1, false, imgOptions, 0, "")
	}
	return pdf.OutputFileAndClose(pdfFileName)
}

func convertPixelToMM(input float64) float64 {
	return 0.2645833333 * input
}

func getPageSize(h, w float64) (string, float64, float64) {
	var heightLimit float64
	if h > w {
		//dealing with potrait
		heightLimit = 210
	} else {
		//dealing with landscape
		heightLimit = 297
	}
	if h <= heightLimit {
		log.Println("it is a A4 size")
		return "A4", 297, 210
	} else {
		log.Println("it is a A3 size")
		return "A3", 594, 420
	}
}

func getImageDimensions(imgFileName string) (int, int, error) {
	imgFileNameSplit := strings.Split(imgFileName, ".")
	imgFileType := strings.ToLower(imgFileNameSplit[len(imgFileNameSplit)-1])
	if imgFileType != "jpg" && imgFileType != "jpeg" && imgFileType != "png" {
		errorMessage := fmt.Sprintf("%s image type unsupported", imgFileType)
		log.Println(errorMessage)
		return -1, -1, errors.New(errorMessage)
	}
	imgFile, imgFileOpenErr := os.Open(imgFileName)
	if imgFileOpenErr != nil {
		log.Println("error while opening image file")
		return -1, -1, imgFileOpenErr
	}
	imgConfig, _, imgDecodeErr := image.DecodeConfig(imgFile)
	if imgDecodeErr != nil {
		log.Printf("error while determining image dimensions for %s\n", imgFileName)
		return -1, -1, imgDecodeErr
	}
	log.Printf("successfully fetched image dimensions for %s - height: %d, width: %d\n", imgFileName, imgConfig.Height, imgConfig.Width)
	return imgConfig.Height, imgConfig.Width, nil
}
