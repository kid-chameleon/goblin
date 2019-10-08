package models

var (
	// StatusEndpoint is the API endpoint for status
	StatusEndpoint = "/status"
)

// Status is the top level Status response
type Status struct {
	Data AssetList `json:"data"`
}
