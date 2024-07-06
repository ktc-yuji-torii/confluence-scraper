package parser

import (
	"encoding/json"

	"github.com/ktc-yuji-torii/confluence-scraper/models"
)

// ParseSpaceData parses JSON data and returns a list of Space objects.
func ParseSpaceData(jsonData string) ([]models.Space, error) {
	var response struct {
		Results []models.Space `json:"results"`
		Links   struct {
			Next string `json:"next"`
			Base string `json:"base"`
		} `json:"_links"`
	}

	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

// ParseSingleSpaceData parses JSON data and returns a single Space object.
func ParseSingleSpaceData(jsonData string) (models.Space, error) {
	var space models.Space

	err := json.Unmarshal([]byte(jsonData), &space)
	if err != nil {
		return space, err
	}

	return space, nil
}
