package services

import (
	"errors"
	"image"
	"net/http"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/internals/api/dto"
)

func ProcessorHandler(w http.ResponseWriter, r *http.Request) {
	requestData, statusCode, errorMsg := parseProcessorRequest(r)
	if statusCode != http.StatusOK {
		http.Error(w, errorMsg, statusCode)
	}

	rgbaImage := convertImageToRGBA(requestData.Img)
	engines.ApplyFilterPipeline(rgbaImage, &requestData.Pipeline, requestData.IsConcurrent)
	encodeResponseImage(w, rgbaImage, requestData.ImgFormat)
}

func parseProcessorRequest(r *http.Request) (*dto.ProcessorRequestDTO, int, string) {
	err := r.ParseMultipartForm(15 << 20)
	if err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			return nil, http.StatusRequestEntityTooLarge, "Request body is too large"
		}
		return nil, http.StatusBadRequest, "Malformed multipart/form-data request"
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		return nil, http.StatusBadRequest, "No image provided"
	}
	defer file.Close()

	filterName := r.URL.Query().Get("filter")
	if filterName == "" {
		return nil, http.StatusBadRequest, "Parameter \"filter\" is required"
	}

	filter, err := filters.GetFilterPipeline(filterName)
	if err != nil {
		return nil, http.StatusNotFound, "Requested filter does not exist"
	}

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, http.StatusBadRequest, "Image format is not supported"
	}

	isConcurrent := false
	if c := r.URL.Query().Get("concurrent"); c == "true" {
		isConcurrent = true
	}

	return &dto.ProcessorRequestDTO{
		Img:          img,
		ImgFormat:    format,
		Pipeline:     filter,
		IsConcurrent: isConcurrent,
	}, http.StatusOK, ""
}
