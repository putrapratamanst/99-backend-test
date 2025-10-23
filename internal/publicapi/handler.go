package publicapi

import (
	"99-backend-exercise/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler defines the HTTP handler for public API endpoints
type Handler struct {
	publicAPIService Service
}

// NewHandler creates a new public API handler
func NewHandler(publicAPIService Service) *Handler {
	return &Handler{
		publicAPIService: publicAPIService,
	}
}

// PublicListingsRequest represents the request for public listings
type PublicListingsRequest struct {
	PageNum  int  `form:"page_num" json:"page_num"`
	PageSize int  `form:"page_size" json:"page_size"`
	UserID   *int `form:"user_id" json:"user_id,omitempty"`
}

// CreateUserRequest represents the request for creating a user via public API
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateListingRequest represents the request for creating a listing via public API
type CreateListingRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	ListingType string `json:"listing_type" binding:"required,oneof=rent sale"`
	Price       int    `json:"price" binding:"required,min=1"`
}

// GetListings handles GET /public-api/listings
func (h *Handler) GetListings(c *gin.Context) {
	var request PublicListingsRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	// Set defaults
	if request.PageNum <= 0 {
		request.PageNum = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = 10
	}

	listings, err := h.publicAPIService.GetListings(request.PageNum, request.PageSize, request.UserID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get listings", err)
		return
	}

	response := map[string]interface{}{
		"listings": listings,
	}
	utils.RespondWithSuccess(c, response)
}

// CreateUser handles POST /public-api/users
func (h *Handler) CreateUser(c *gin.Context) {
	var request CreateUserRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	user, err := h.publicAPIService.CreateUser(request.Name)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	// Return user directly (as per API spec)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

// CreateListing handles POST /public-api/listings
func (h *Handler) CreateListing(c *gin.Context) {
	var request CreateListingRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	listing, err := h.publicAPIService.CreateListing(request.UserID, request.ListingType, request.Price)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create listing", err)
		return
	}

	// Return listing directly (as per API spec)
	c.JSON(http.StatusOK, map[string]interface{}{
		"listing": listing,
	})
}
