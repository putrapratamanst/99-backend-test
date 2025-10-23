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
type HTTPClient interface {
	Get(url string) (*http.Response, error)
	PostForm(url string, data url.Values) (*http.Response, error)
}
type DefaultHTTPClient struct {
	client *http.Client
}
func NewHTTPClient() HTTPClient {
	return &DefaultHTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}
func (c *DefaultHTTPClient) PostForm(url string, data url.Values) (*http.Response, error) {
	return c.client.PostForm(url, data)
}
type ServiceClient struct {
	httpClient        HTTPClient
	userServiceURL    string
	listingServiceURL string
}
func NewServiceClient(userServiceURL, listingServiceURL string) *ServiceClient {
	return &ServiceClient{
		httpClient:        NewHTTPClient(),
		userServiceURL:    userServiceURL,
		listingServiceURL: listingServiceURL,
	}
}
type UserServiceResponse struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}
type ListingServiceResponse struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}
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
	if data, ok := response.Data.(map[string]interface{}); ok {
		if user, ok := data["user"].(map[string]interface{}); ok {
			return user, nil
		}
	}
	return nil, fmt.Errorf("unexpected response format")
}
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
	if data, ok := response.Data.(map[string]interface{}); ok {
		if user, ok := data["user"].(map[string]interface{}); ok {
			return user, nil
		}
	}
	return nil, fmt.Errorf("unexpected response format")
}
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
	if data, ok := response.Data.(map[string]interface{}); ok {
		if listings, ok := data["listings"].([]interface{}); ok {
			return listings, nil
		}
	}
	return nil, fmt.Errorf("unexpected response format")
}
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
	if data, ok := response.Data.(map[string]interface{}); ok {
		if listing, ok := data["listing"].(map[string]interface{}); ok {
			return listing, nil
		}
	}
	return nil, fmt.Errorf("unexpected response format")
}
