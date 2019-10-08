package models

// Asset is a naive Type/ID structure used in many places,
// sometimes it is empty
type Asset struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

// AssetList is a naive list of Assets
type AssetList struct {
	Data []Asset `json:"data,omitempty"`
}

// APIResponse is a top-level api response container
type APIResponse struct {
	Links `json:"links"`
	Meta  `json:"meta,omitempty"`
}

type Links struct {
	Self string `json:"self"`
}

type LinksWithSchema struct {
	Links

	Schema string `json:"schema,omitempty"`
}

type Meta struct{}

type Tags interface{}
