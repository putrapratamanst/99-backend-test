package publicapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// HTTPClient defines the interface for HTTP operations
type HTTPClient interface {
	Get(url string) (*http.Response, error)
	PostForm(url string, data url.Values) (*http.Response, error)
}

// DefaultHTTPClient is the default implementation
type DefaultHTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient() HTTPClient {
	return &DefaultHTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Get performs a GET request
func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

// PostForm performs a POST request with form data
func (c *DefaultHTTPClient) PostForm(url string, data url.Values) (*http.Response, error) {
	return c.client.PostForm(url, data)
}

// ServiceClient handles communication with microservices
type ServiceClient struct {
	httpClient        HTTPClient
	userServiceURL    string
	listingServiceURL string
}

// NewServiceClient creates a new service client
func NewServiceClient(userServiceURL, listingServiceURL string) *ServiceClient {
	return &ServiceClient{
		httpClient:        NewHTTPClient(),
		userServiceURL:    userServiceURL,
		listingServiceURL: listingServiceURL,
	}
}

// UserServiceResponse represents the response from user service
type UserServiceResponse struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

// ListingServiceResponse represents the response from listing service
type ListingServiceResponse struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

// GetUser fetches a user by ID from user service
func (sc *ServiceClient) GetUser(userID int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/users/%d", sc.userServiceURL, userID)
	resp, err := sc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response UserServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("user service error: %v", response.Error)
	}

	// Extract user data from the response
	if data, ok := response.Data.(map[string]interface{}); ok {
		if user, ok := data["user"].(map[string]interface{}); ok {
			return user, nil
		}
	}

	return nil, fmt.Errorf("unexpected response format")
}

// CreateUser creates a user via user service
func (sc *ServiceClient) CreateUser(name string) (map[string]interface{}, error) {
	data := url.Values{
		"name": {name},
	}

	url := fmt.Sprintf("%s/users", sc.userServiceURL)
	resp, err := sc.httpClient.PostForm(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response UserServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("user service error: %v", response.Error)
	}

	// Extract user data from the response
	if data, ok := response.Data.(map[string]interface{}); ok {
		if user, ok := data["user"].(map[string]interface{}); ok {
			return user, nil
		}
	}

	return nil, fmt.Errorf("unexpected response format")
}

// GetListings fetches listings from listing service
func (sc *ServiceClient) GetListings(pageNum, pageSize int, userID *int) ([]interface{}, error) {
	params := url.Values{
		"page_num":  {strconv.Itoa(pageNum)},
		"page_size": {strconv.Itoa(pageSize)},
	}

	if userID != nil {
		params.Set("user_id", strconv.Itoa(*userID))
	}

	url := fmt.Sprintf("%s/listings?%s", sc.listingServiceURL, params.Encode())
	resp, err := sc.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call listing service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response ListingServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("listing service error: %v", response.Error)
	}

	// Extract listings data from the response
	if data, ok := response.Data.(map[string]interface{}); ok {
		if listings, ok := data["listings"].([]interface{}); ok {
			return listings, nil
		}
	}

	return nil, fmt.Errorf("unexpected response format")
}

// CreateListing creates a listing via listing service
func (sc *ServiceClient) CreateListing(userID int, listingType string, price int) (map[string]interface{}, error) {
	data := url.Values{
		"user_id":      {strconv.Itoa(userID)},
		"listing_type": {listingType},
		"price":        {strconv.Itoa(price)},
	}

	url := fmt.Sprintf("%s/listings", sc.listingServiceURL)
	resp, err := sc.httpClient.PostForm(url, data)
	if err != nil {
		return nil, fmt.Errorf("failed to call listing service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response ListingServiceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !response.Result {
		return nil, fmt.Errorf("listing service error: %v", response.Error)
	}

	// Extract listing data from the response
	if data, ok := response.Data.(map[string]interface{}); ok {
		if listing, ok := data["listing"].(map[string]interface{}); ok {
			return listing, nil
		}
	}

	return nil, fmt.Errorf("unexpected response format")
}
