package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// BaseURL is the root endpoint for the Nexon OpenAPI.
	BaseURL = "https://open.api.nexon.com/"
	// DefaultTimeout matches the 5s default used in other language bindings.
	DefaultTimeout = 5 * time.Second
)

// Client wraps the shared HTTP configuration used by region-specific MapleStory clients.
type Client struct {
	apiKey     string
	subURL     string
	timezone   *time.Location
	httpClient *http.Client
	baseURL    string
}

// NewClient constructs a Client with the provided API key, sub-route, and timezone.
func NewClient(apiKey string, subURL string, timezone *time.Location) *Client {
	if timezone == nil {
		timezone = time.UTC
	}

	return &Client{
		apiKey:     apiKey,
		subURL:     strings.Trim(subURL, "/"),
		timezone:   timezone,
		baseURL:    BaseURL,
		httpClient: &http.Client{Timeout: DefaultTimeout},
	}
}

// Timezone returns the client's timezone.
func (c *Client) Timezone() *time.Location {
	return c.timezone
}

// SubURL returns the region-specific sub-path (e.g., "maplestory", "maplestorytw").
func (c *Client) SubURL() string {
	return c.subURL
}

// Base returns the base URL used for API calls.
func (c *Client) Base() string {
	return c.baseURL
}

// SetBase overrides the base URL (useful for testing or future mock servers).
func (c *Client) SetBase(base string) {
	c.baseURL = strings.TrimRight(base, "/") + "/"
}

// Timeout returns the configured HTTP timeout.
func (c *Client) Timeout() time.Duration {
	return c.httpClient.Timeout
}

// SetTimeout adjusts the HTTP timeout.
func (c *Client) SetTimeout(d time.Duration) {
	c.httpClient.Timeout = d
}

// Fetch executes a GET request, injecting the API key header and handling MapleStory error envelopes.
// Query parameters with empty values are omitted. If target is non-nil, the response is unmarshaled into it.
func (c *Client) Fetch(ctx context.Context, path string, query map[string]string, target any) error {
	if ctx == nil {
		ctx = context.Background()
	}

	fullURL, err := c.buildURL(path, query)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-nxopen-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Attempt to decode the common error envelope before doing anything else.
	var errBody MapleStoryErrorBody
	if err := json.Unmarshal(data, &errBody); err == nil && errBody.Error.Name != "" {
		return newAPIError(errBody.Error)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unexpected status %d from MapleStory API", resp.StatusCode)
	}

	if target != nil {
		if err := json.Unmarshal(data, target); err != nil {
			return err
		}
	}

	return nil
}

// ToDateString converts a date-like value into the API's expected string using the client's timezone.
func (c *Client) ToDateString(value any, min *Date) (string, error) {
	return ToDateString(value, c.timezone, min)
}

// ProperDefaultDate returns the latest available date given an update window using the client's timezone.
func (c *Client) ProperDefaultDate(opts LatestAPIUpdateTimeOptions) Date {
	return GetProperDefaultDate(time.Now(), c.timezone, opts)
}

func (c *Client) buildURL(path string, query map[string]string) (string, error) {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	resolvedPath := strings.TrimLeft(path, "/")
	if c.subURL != "" && !strings.HasPrefix(resolvedPath, c.subURL) {
		resolvedPath = c.subURL + "/" + resolvedPath
	}
	base.Path = resolvedPath

	q := base.Query()
	for key, value := range query {
		if strings.TrimSpace(value) == "" {
			continue
		}
		q.Set(key, value)
	}
	base.RawQuery = q.Encode()

	return base.String(), nil
}
