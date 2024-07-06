package models

type OutputPage struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

// Page represents the information of a Confluence page.
type Page struct {
	ID                       string         `json:"id"`                       // The ID of the page
	Status                   string         `json:"status"`                   // The status of the page (e.g., current)
	Title                    string         `json:"title"`                    // The title of the page
	SpaceID                  string         `json:"spaceId"`                  // The ID of the space to which the page belongs
	ParentID                 string         `json:"parentId"`                 // The ID of the parent page
	ParentType               string         `json:"parentType"`               // The type of the parent (e.g., page)
	Position                 int            `json:"position"`                 // The position of the page
	AuthorID                 string         `json:"authorId"`                 // The ID of the author of the page
	OwnerID                  string         `json:"ownerId"`                  // The ID of the owner of the page
	LastOwnerID              string         `json:"lastOwnerId"`              // The ID of the last owner of the page
	CreatedAt                string         `json:"createdAt"`                // The creation date of the page
	Version                  PageVersion    `json:"version"`                  // The version information of the page
	Body                     PageBody       `json:"body"`                     // The body content of the page
	Labels                   PageLabels     `json:"labels"`                   // The labels associated with the page
	Properties               PageProperties `json:"properties"`               // The properties of the page
	Operations               PageOperations `json:"operations"`               // The operations available for the page
	Likes                    PageLikes      `json:"likes"`                    // The likes on the page
	Versions                 PageVersions   `json:"versions"`                 // The versions of the page
	IsFavoritedByCurrentUser bool           `json:"isFavoritedByCurrentUser"` // Whether the page is favorited by the current user
	Links                    PageLinks      `json:"_links"`                   // The links related to the page
	Content                  string         `json:"content"`                  // The plain text content of the page
	URL                      string         `json:"url"`                      // The URL of the page
}

// PageVersion represents the version information of a Confluence page.
type PageVersion struct {
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	Number    int    `json:"number"`
	MinorEdit bool   `json:"minorEdit"`
	AuthorID  string `json:"authorId"`
}

// PageBody represents the body content of a Confluence page.
type PageBody struct {
	Storage struct {
		Value string `json:"value"`
	} `json:"storage"`
	AtlasDocFormat struct {
		Value string `json:"value"`
	} `json:"atlas_doc_format"`
	View struct {
		Value string `json:"value"`
	} `json:"view"`
}

// PageLabels represents the labels associated with a Confluence page.
type PageLabels struct {
	Results []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Prefix string `json:"prefix"`
	} `json:"results"`
	Meta struct {
		HasMore bool   `json:"hasMore"`
		Cursor  string `json:"cursor"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

// PageProperties represents the properties of a Confluence page.
type PageProperties struct {
	Results []struct {
		ID      string `json:"id"`
		Key     string `json:"key"`
		Version struct {
			// Add fields as necessary
		} `json:"version"`
	} `json:"results"`
	Meta struct {
		HasMore bool   `json:"hasMore"`
		Cursor  string `json:"cursor"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

// PageOperations represents the operations available for a Confluence page.
type PageOperations struct {
	Results []struct {
		Operation  string `json:"operation"`
		TargetType string `json:"targetType"`
	} `json:"results"`
	Meta struct {
		HasMore bool   `json:"hasMore"`
		Cursor  string `json:"cursor"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

// PageLikes represents the likes on a Confluence page.
type PageLikes struct {
	Results []struct {
		AccountID string `json:"accountId"`
	} `json:"results"`
	Meta struct {
		HasMore bool   `json:"hasMore"`
		Cursor  string `json:"cursor"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

// PageVersions represents the versions of a Confluence page.
type PageVersions struct {
	Results []struct {
		CreatedAt string `json:"createdAt"`
		Message   string `json:"message"`
		Number    int    `json:"number"`
		MinorEdit bool   `json:"minorEdit"`
		AuthorID  string `json:"authorId"`
	} `json:"results"`
	Meta struct {
		HasMore bool   `json:"hasMore"`
		Cursor  string `json:"cursor"`
	} `json:"meta"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}

// PageMeta represents metadata for various components of a Confluence page.
type PageMeta struct {
	HasMore bool   `json:"hasMore"`
	Cursor  string `json:"cursor"`
}

// PageLinks represents the links related to various components of a Confluence page.
type PageLinks struct {
	Self string `json:"self"`
	Base string `json:"base"`
}

// PageContent represents the content of a Confluence page.
type PageContent struct {
	ID      string `json:"id"`      // The ID of the page
	Title   string `json:"title"`   // The title of the page
	Content string `json:"content"` // The content of the page (in HTML format)
	Space   Space  `json:"space"`   // The space to which the page belongs
}

// Space represents the information of a Confluence space.
type Space struct {
	ID   string `json:"id"`   // The ID of the space
	Key  string `json:"key"`  // The key of the space
	Name string `json:"name"` // The name of the space
}
