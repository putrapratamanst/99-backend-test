package publicapi

import (
	"99-backend-exercise/internal/models"
	"fmt"
)

// Service defines the interface for public API business logic
type Service interface {
	GetListings(pageNum, pageSize int, userID *int) ([]models.PublicListingResponse, error)
	CreateUser(name string) (map[string]interface{}, error)
	CreateListing(userID int, listingType string, price int) (map[string]interface{}, error)
}

// service is the concrete implementation of Service interface
type service struct {
	serviceClient *ServiceClient
}

// NewService creates a new public API service
func NewService(serviceClient *ServiceClient) Service {
	return &service{
		serviceClient: serviceClient,
	}
}

// GetListings retrieves listings with user details
func (s *service) GetListings(pageNum, pageSize int, userID *int) ([]models.PublicListingResponse, error) {
	// Get listings from listing service
	listings, err := s.serviceClient.GetListings(pageNum, pageSize, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get listings: %w", err)
	}

	var result []models.PublicListingResponse

	// For each listing, get the user details
	for _, listingData := range listings {
		listing, ok := listingData.(map[string]interface{})
		if !ok {
			continue
		}

		// Extract user_id from listing
		userIDFloat, ok := listing["user_id"].(float64)
		if !ok {
			continue
		}
		listingUserID := int(userIDFloat)

		// Get user details
		user, err := s.serviceClient.GetUser(listingUserID)
		if err != nil {
			// If user not found, skip this listing or log error
			continue
		}

		// Build public listing response
		publicListing := models.PublicListingResponse{
			ID:          int(listing["id"].(float64)),
			ListingType: listing["listing_type"].(string),
			Price:       int(listing["price"].(float64)),
			CreatedAt:   int64(listing["created_at"].(float64)),
			UpdatedAt:   int64(listing["updated_at"].(float64)),
			User: models.UserResponse{
				ID:        int(user["id"].(float64)),
				Name:      user["name"].(string),
				CreatedAt: int64(user["created_at"].(float64)),
				UpdatedAt: int64(user["updated_at"].(float64)),
			},
		}

		result = append(result, publicListing)
	}

	return result, nil
}

// CreateUser creates a new user via user service
func (s *service) CreateUser(name string) (map[string]interface{}, error) {
	return s.serviceClient.CreateUser(name)
}

// CreateListing creates a new listing via listing service
func (s *service) CreateListing(userID int, listingType string, price int) (map[string]interface{}, error) {
	return s.serviceClient.CreateListing(userID, listingType, price)
}
