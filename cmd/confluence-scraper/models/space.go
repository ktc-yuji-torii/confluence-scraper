package models

// Space represents the information of a Confluence space.
type Space struct {
	ID             string `json:"id"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	AuthorID       string `json:"authorId"`
	CreatedAt      string `json:"createdAt"`
	HomepageID     string `json:"homepageId"`
	Description    Description `json:"description"`
	Icon           Icon `json:"icon"`
	WebUI          string `json:"webui"`
}

type Description struct {
	Plain struct {
		Value string `json:"value"`
	} `json:"plain"`
	View struct {
		Value string `json:"value"`
	} `json:"view"`
}

type Icon struct {
	Path           string `json:"path"`
	APIDownloadLink string `json:"apiDownloadLink"`
}
