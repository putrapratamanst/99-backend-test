package publicapi
import (
	"99-backend-exercise/internal/models"
	"fmt"
)
type Service interface {
	GetListings(pageNum, pageSize int, userID *int) ([]models.PublicListingResponse, error)
	CreateUser(name string) (map[string]interface{}, error)
	CreateListing(userID int, listingType string, price int) (map[string]interface{}, error)
}
type service struct {
	serviceClient *ServiceClient
}
func NewService(serviceClient *ServiceClient) Service {
	return &service{
		serviceClient: serviceClient,
	}
}
func (s *service) GetListings(pageNum, pageSize int, userID *int) ([]models.PublicListingResponse, error) {
	listings, err := s.serviceClient.GetListings(pageNum, pageSize, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get listings: %w", err)
	}
	var result []models.PublicListingResponse
	for _, listingData := range listings {
		listing, ok := listingData.(map[string]interface{})
		if !ok {
			continue
		}
		userIDFloat, ok := listing["user_id"].(float64)
		if !ok {
			continue
		}
		listingUserID := int(userIDFloat)
		user, err := s.serviceClient.GetUser(listingUserID)
		if err != nil {
			continue
		}
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
func (s *service) CreateUser(name string) (map[string]interface{}, error) {
	return s.serviceClient.CreateUser(name)
}
func (s *service) CreateListing(userID int, listingType string, price int) (map[string]interface{}, error) {
	return s.serviceClient.CreateListing(userID, listingType, price)
}
