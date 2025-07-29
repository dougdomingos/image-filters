package services

import (
	"net/http"

	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/internals/api/dto"
)

// ListFilterHandler provides a service to list all avaliable filters in the
// API.
func ListFilterHandler(w http.ResponseWriter, r *http.Request) {
	filterIDs := filters.GetAvaliableFilterIDs()
	response := dto.ListFiltersResponse{Filters: filterIDs}
	sendJSONResponse(w, response, http.StatusOK)
}
