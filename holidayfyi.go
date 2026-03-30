// Package holidayfyi provides a Go client for the HolidayFYI API.
//
// HolidayFYI (holidayfyi.com) — Holiday dates, Easter calculation, and 200+ country calendars.
// This client requires no authentication and has zero external dependencies.
//
// Usage:
//
//	client := holidayfyi.NewClient()
//	result, err := client.Search("example")
package holidayfyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the HolidayFYI API.
const DefaultBaseURL = "https://holidayfyi.com/api"

// Client is a HolidayFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new HolidayFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("holidayfyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("holidayfyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("holidayfyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across all content.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Entity returns details for a holiday by slug.
func (c *Client) Entity(slug string) (*EntityDetail, error) {
	var result EntityDetail
	if err := c.get("/holiday/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random holiday.
func (c *Client) Random() (*EntityDetail, error) {
	var result EntityDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
