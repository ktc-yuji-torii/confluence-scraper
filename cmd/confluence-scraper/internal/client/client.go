package client

import (
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"

	"github.com/ktc-yuji-torii/confluence-scraper/config"
	"github.com/ktc-yuji-torii/confluence-scraper/internal/parser"
	"github.com/ktc-yuji-torii/confluence-scraper/models"
)

type ConfluenceClient struct {
	config config.Config
	client *http.Client
	logger *slog.Logger
}

func NewConfluenceClient(cfg config.Config, logger *slog.Logger) *ConfluenceClient {
	return &ConfluenceClient{
		config: cfg,
		client: &http.Client{},
		logger: logger,
	}
}

func (c *ConfluenceClient) GetChildPages(parentPageID string) (string, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/pages/%s/children", c.config.BaseURL, parentPageID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.config.Username, c.config.APIToken)))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch child pages: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *ConfluenceClient) GetPageContent(pageID string) (string, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/pages/%s?body-format=storage", c.config.BaseURL, pageID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Error creating request", "error", err, "url", url)
		return "", err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.config.Username, c.config.APIToken)))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch page content: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *ConfluenceClient) GetChildPagesRecursively(parentPageID string, cfg config.Config) ([]models.OutputPage, error) {
	var allPages []models.OutputPage
	var mu sync.Mutex
	var wg sync.WaitGroup

	c.logger.Debug("Fetching child pages", "parentPageID", parentPageID)

	childPagesJSON, err := c.GetChildPages(parentPageID)
	if err != nil {
		return nil, err
	}

	childPages, err := parser.ParseChildPages(childPagesJSON)
	if err != nil {
		return nil, err
	}

	for _, childPage := range childPages.Results {
		wg.Add(1)
		go func(childPage models.ChildPage) {
			defer wg.Done()

			c.logger.Debug("Fetching content for child page", "childPageID", childPage.ID, "title", childPage.Title)

			pageContentJSON, err := c.GetPageContent(childPage.ID)
			if err != nil {
				c.logger.Error("Error fetching content for page", "pageID", childPage.ID, "error", err)
				return
			}

			page, err := parser.ParsePageContent(pageContentJSON, cfg)
			if err != nil {
				c.logger.Error("Error parsing content for page", "pageID", childPage.ID, "error", err)
				return
			}

			outputPage := models.OutputPage{
				ID:      page.ID,
				Title:   page.Title,
				Content: page.Content,
				URL:     page.URL,
			}

			mu.Lock()
			allPages = append(allPages, outputPage)
			mu.Unlock()

			subPages, err := c.GetChildPagesRecursively(childPage.ID, cfg)
			if err != nil {
				c.logger.Error("Error fetching child pages", "parentPageID", childPage.ID, "error", err)
				return
			}

			mu.Lock()
			allPages = append(allPages, subPages...)
			mu.Unlock()
		}(childPage)
	}

	wg.Wait()

	c.logger.Debug("Completed fetching child pages", "parentPageID", parentPageID)
	return allPages, nil
}
