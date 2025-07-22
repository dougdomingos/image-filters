package api

import (
	"fmt"
	"image"
	"net/http"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
)

func benchmarkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Benchmark route accessed")
}

func processorHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Request is too large or malformed", http.StatusBadRequest)
		return
	}

	// TODO: abstract validation logic into external function
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image was not provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filterName := r.URL.Query().Get("filter")
	if filterName == "" {
		http.Error(w, "Filter was not provided", http.StatusBadRequest)
		return
	}

	filter, err := filters.GetFilterPipeline(filterName)
	if err != nil {
		http.Error(w, "Requested filter does not exist", http.StatusNotFound)
		return
	}

	img, format, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Image format is not supported", http.StatusBadRequest)
		return
	}
	
	isConcurrent := false
	if c := r.URL.Query().Get("concurrent"); c == "true" {
		isConcurrent = true
	}
	
	rgbaImage := convertImageToRGBA(img)
	engines.ApplyFilterPipeline(rgbaImage, &filter, isConcurrent)
	encodeResponseImage(w, rgbaImage, format)
}
