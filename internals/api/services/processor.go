package services

import (
	"encoding/json"
	"errors"
	"image"
	"net/http"
	"os"
	"strings"
	"time"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/internals/api/dto"
	"dougdomingos.com/image-filters/internals/utils"
	"dougdomingos.com/image-filters/pipelines"
)

// ProcessorHandler provides the image processing service to the API. It
// accepts an image through a multipart/form-data request. The filter
// and execution mode are specified through URL parameters.
func ProcessorHandler(w http.ResponseWriter, r *http.Request) {
	requestData, statusCode, errorMsg := parseProcessorRequest(r)
	if statusCode != http.StatusOK {
		http.Error(w, errorMsg, statusCode)
	}

	rgbaImage := convertImageToRGBA(requestData.Img)
	engines.ProcessRecipe(rgbaImage, &requestData.Recipe)

	outputDir := os.Getenv("API_OUTPUT_DIR")
	outputFilename := utils.GetProcessedImageFilename(requestData.ImgFilename, time.Now().Format("20060102_150405"))
	_, err := utils.SaveImage(rgbaImage, requestData.ImgFormat, outputDir, outputFilename)
	if err != nil {
		http.Error(w, "Unable to store image on disk", http.StatusInsufficientStorage)
	}

	response := dto.BuildProcessorResponse(outputFilename)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// parseProcessorRequest extracts and validates the required parameters from the
// HTTP request intended for the processor service. It enforces constraints on
// request size (max. 15 MB), required fields, and image format.
func parseProcessorRequest(r *http.Request) (*dto.ProcessorRequestDTO, int, string) {
	err := r.ParseMultipartForm(15 << 20)
	if err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			return nil, http.StatusRequestEntityTooLarge, "Request body is too large"
		}
		return nil, http.StatusBadRequest, "Malformed multipart/form-data request"
	}
	
	file, header, err := r.FormFile("image")
	if err != nil {
		return nil, http.StatusBadRequest, "No image provided"
	}
	defer file.Close()

	filters := r.URL.Query().Get("filters")
	if filters == "" {
		return nil, http.StatusBadRequest, "Parameter \"filters\" is required"
	}

	recipe, err := pipelines.BuildRecipe(strings.Split(filters, ","))
	if err != nil {
		return nil, http.StatusNotFound, "Requested filter does not exist"
	}

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, http.StatusBadRequest, "Image format is not supported"
	}

	return &dto.ProcessorRequestDTO{
		Img:         img,
		ImgFormat:   format,
		ImgFilename: header.Filename,
		Recipe:      recipe,
	}, http.StatusOK, ""
}
