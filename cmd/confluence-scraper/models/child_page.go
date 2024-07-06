package models

// ChildPage represents a child page in Confluence.
type ChildPage struct {
	Result
}

// Result represents the individual result item in the JSON.
type Result struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	SpaceID       string `json:"spaceId"`
	ChildPosition int    `json:"childPosition"`
}

// ChildPages represents a list of child pages in Confluence.
type ChildPages struct {
	Results []ChildPage `json:"results"` // List of child pages
}

// ChildPagesResponse represents the entire JSON response.
type ChildPagesResponse struct {
	Results []Result `json:"results"`
	Links   Links    `json:"_links"`
}

// Links represents the _links part of the JSON response.
type Links struct {
	Next string `json:"next"`
	Base string `json:"base"`
}
