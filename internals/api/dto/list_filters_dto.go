package dto

// ListFiltersResponse represents the JSON response for listing the avaliable
// filters in the API.
type ListFiltersResponse struct {
	Filters []string `json:"filters"`
}