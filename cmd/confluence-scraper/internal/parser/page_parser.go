package parser

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strconv"
	"strings"

	"github.com/ktc-yuji-torii/confluence-scraper/config"
	"github.com/ktc-yuji-torii/confluence-scraper/models"
)

// Unicodeエスケープシーケンスをデコードする関数
func decodeUnicodeEscapeSequences(input string) string {
	re := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		r, _ := strconv.ParseInt(match[2:], 16, 32)
		return string(rune(r))
	})
}

var (
	reNewLineTags  = regexp.MustCompile(`</?(h[1-6]|p|div|li|br|tr|table|thead|tbody|tfoot|caption|dl|dt|dd|figure|figcaption|pre|hr)>`)
	reListTags     = regexp.MustCompile(`</?(ul|ol)>`)
	reListItemTags = regexp.MustCompile(`</?li>`)
	reHeaders      = regexp.MustCompile(`<(h[1-6])>(.*?)</h[1-6]>`)
	reBoldTags     = regexp.MustCompile(`</?b>|</?strong>`)
	reItalicTags   = regexp.MustCompile(`</?i>|</?em>`)
	reTags         = regexp.MustCompile(`<.*?>`)
	reSpaces       = regexp.MustCompile(`\s+`)
)

func removeHTMLTagsAndFormat(input string) string {
	decoded := html.UnescapeString(input)
	decoded = decodeUnicodeEscapeSequences(decoded)
	formatted := reNewLineTags.ReplaceAllString(decoded, "\n")
	formatted = reListTags.ReplaceAllString(formatted, "\n\n")
	formatted = reListItemTags.ReplaceAllStringFunc(formatted, func(m string) string {
		return "- "
	})
	formatted = reHeaders.ReplaceAllStringFunc(formatted, func(m string) string {
		header := reHeaders.FindStringSubmatch(m)
		level := header[1][1:] // Extract the header level
		content := header[2]
		return fmt.Sprintf("%s %s", strings.Repeat("#", int(level[0]-'0')), content)
	})
	formatted = reBoldTags.ReplaceAllString(formatted, "**")
	formatted = reItalicTags.ReplaceAllString(formatted, "*")
	formatted = reTags.ReplaceAllString(formatted, "")
	formatted = reSpaces.ReplaceAllString(formatted, " ")
	return strings.TrimSpace(formatted)
}

// ParsePageContent parses JSON data and returns a Page object.
func ParsePageContent(jsonData string, cfg config.Config) (models.Page, error) {
	var pageData models.Page

	err := json.Unmarshal([]byte(jsonData), &pageData)
	if err != nil {
		return models.Page{}, err
	}

	// HTMLエスケープシーケンスをデコードしてContentを作成
	decodedContent := html.UnescapeString(pageData.Body.Storage.Value)
	plainTextContent := removeHTMLTagsAndFormat(decodedContent)
	pageData.Content = plainTextContent
	pageData.URL = fmt.Sprintf("%s/wiki/spaces/%s/pages/%s", cfg.BaseURL, pageData.SpaceID, pageData.ID)
	return pageData, nil
}

func ParseChildPages(jsonData string) (models.ChildPages, error) {
	var childPages models.ChildPages
	err := json.Unmarshal([]byte(jsonData), &childPages)
	if err != nil {
		return childPages, err
	}
	return childPages, nil
}

// ConvertToOutputPage converts a Page object to an OutputPage object.
func ConvertToOutputPage(page models.Page, cfg config.Config) models.OutputPage {
	decodedContent := html.UnescapeString(page.Body.Storage.Value)
	plainTextContent := removeHTMLTagsAndFormat(decodedContent)

	return models.OutputPage{
		ID:      page.ID,
		Title:   page.Title,
		Content: plainTextContent,
		URL:     fmt.Sprintf("%s/wiki/spaces/%s/pages/%s", cfg.BaseURL, page.SpaceID, page.ID),
	}
}

// ConvertPagesToOutputPages converts a slice of Page objects to a slice of OutputPage objects.
func ConvertPagesToOutputPages(pages []models.Page, cfg config.Config) []models.OutputPage {
	var outputPages []models.OutputPage
	for _, page := range pages {
		outputPage := ConvertToOutputPage(page, cfg)
		outputPages = append(outputPages, outputPage)
	}
	return outputPages
}
