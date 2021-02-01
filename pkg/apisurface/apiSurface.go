package apisurface

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/google/uuid"
	"github.com/MadhavanKR/pdfWorker/pkg/pdfservices"
)

func ConvertImagesToPdfHandler(w http.ResponseWriter, r *http.Request) {
	requestId, rqstIdGenErr := uuid.NewRandom()
	if rqstIdGenErr != nil {
		errorResponse(w, rqstIdGenErr, http.StatusInternalServerError)
		return
	}
	mpFormParseErr := r.ParseMultipartForm(10 * 1024 * 1024) //limiting file size to 10mb
	if mpFormParseErr != nil {
		errorResponse(w, mpFormParseErr, http.StatusInternalServerError)
	}
	uploadedFileMap := r.MultipartForm.File
	imgPaths := make([]string, 0)
	tmpFilesDir := "E:\\uploads\\" + requestId.String()
	os.Mkdir(tmpFilesDir, 0644)
	//defer deleteAllFiles(tmpFilesDir)
	for _, fileHeader := range uploadedFileMap {
		openFile, openFileErr := fileHeader[0].Open()
		log.Printf("fileName: %s - size:%d\n", fileHeader[0].Filename, fileHeader[0].Size)
		if openFileErr != nil {
			log.Printf("error while opening file to download: %v\n",openFileErr)
			errorResponse(w, openFileErr, http.StatusInternalServerError)
		}
		tempFileName := wrapInQuotes(tmpFilesDir + "\\" + fileHeader[0].Filename)
		tempFile, tempFileOpenErr := os.Create(tempFileName)
		if tempFileOpenErr != nil {
			log.Printf("error while opening temp file: %v\n", tempFileOpenErr)
			errorResponse(w, tempFileOpenErr, http.StatusInternalServerError)
			return
		}
		_, fileCopyErr := io.Copy(tempFile, openFile)
		if fileCopyErr != nil {
			log.Printf("error while copying file to temp directory: %v\n",fileCopyErr)
			errorResponse(w, tempFileOpenErr, http.StatusInternalServerError)
			return
		}
		imgPaths = append(imgPaths, tempFileName)
		openFile.Close()
		tempFile.Close()
	}
	log.Printf("successfully copied %d files to %s\n", len(imgPaths), tmpFilesDir)
	log.Println("combining images and converting to pdf")
	outputPdfFilePath := tmpFilesDir + "\\output.pdf"
	log.Printf("pdf will be created at %s\n", outputPdfFilePath)
	pdfConversionErr := pdfservices.ConvertToPdf(outputPdfFilePath, imgPaths)
	if pdfConversionErr != nil {
		errorResponse(w, pdfConversionErr, http.StatusInternalServerError)
	} else {
		headers := make(map[string]string)
		headers["Content-Type"] = "application/octet-stream"
		headers["Content-Disposition"] = fmt.Sprintf("attachment; filename=output.pdf")
		outputPdfFile, opPdfFileErr := os.Open(outputPdfFilePath)
		if opPdfFileErr != nil {
			errorResponse(w, opPdfFileErr, http.StatusInternalServerError)
		} else {
			defer outputPdfFile.Close()
			writeResponse(w, outputPdfFile, headers, http.StatusOK)
		}
	}
}

func errorResponse(w http.ResponseWriter, err error, status int) {
	errorMessage := fmt.Sprintf("error while processing request: %v", err)
	log.Println(errorMessage, err)
	errorResponse := ApiErrorResponse{
		Error:  errorMessage,
		Status: http.StatusBadRequest,
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	writeResponse(w, errorResponse, headers, status)
}

func writeResponse(w http.ResponseWriter, content interface{}, headers map[string]string, status int) {
	responseBody, responseMarshalErr := json.Marshal(content)
	if responseMarshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error. error while forming response"))
	} else {
		if headers!= nil {
			for headerKey, headerValue := range headers {
				w.Header().Add(headerKey, headerValue)
			}
		}
		w.WriteHeader(status)
		w.Write(responseBody)
	}

}

func deleteAllFiles(dirPath string) {
	os.RemoveAll(dirPath)
}

func wrapInQuotes(input string) string {
	//return "\"" + input + "\""
	return input
}
