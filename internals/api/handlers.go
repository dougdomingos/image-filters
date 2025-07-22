package api

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
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

	filterName := r.FormValue("filter")
	if filterName == "" {
		http.Error(w, "Filter was not provided", http.StatusBadRequest)
		return
	}

	filter, err := filters.GetFilterPipeline(filterName)
	if err != nil {
		http.Error(w, "Requested filter does not exist", http.StatusNotFound)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Image format is not supported", http.StatusBadRequest)
		return
	}

	// TODO: check if format conversion is necessary for all image formats
	bounds := img.Bounds()
	rgbaImage := image.NewRGBA(bounds)
	draw.Draw(rgbaImage, bounds, img, bounds.Min, draw.Src)

	// TODO: allow users to select serial or concurrent execution modes
	engines.ApplyFilterPipeline(rgbaImage, &filter, false)

	// TODO: detect image original encoding
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, rgbaImage)
}
