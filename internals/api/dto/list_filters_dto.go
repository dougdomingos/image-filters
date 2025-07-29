package dto

// ListFiltersResponse represents the JSON response for listing the avaliable
// filters in the API.
type ListFiltersResponse struct {

	// Filters contains the slice of all avaliable filter IDs.
	Filters []string `json:"filters"`
}