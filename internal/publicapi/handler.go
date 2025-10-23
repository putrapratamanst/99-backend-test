package publicapi
import (
	"99-backend-exercise/pkg/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)
type Handler struct {
	publicAPIService Service
}
func NewHandler(publicAPIService Service) *Handler {
	return &Handler{
		publicAPIService: publicAPIService,
	}
}
type PublicListingsRequest struct {
	PageNum  int  `form:"page_num" json:"page_num"`
	PageSize int  `form:"page_size" json:"page_size"`
	UserID   *int `form:"user_id" json:"user_id,omitempty"`
}
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
type CreateListingRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	ListingType string `json:"listing_type" binding:"required,oneof=rent sale"`
	Price       int    `json:"price" binding:"required,min=1"`
}
func (h *Handler) GetListings(c *gin.Context) {
	var request PublicListingsRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}
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
func (h *Handler) CreateUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}
	user, err := h.publicAPIService.CreateUser(request.Name)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
func (h *Handler) CreateListing(c *gin.Context) {
	var request CreateListingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}
	listing, err := h.publicAPIService.CreateListing(request.UserID, request.ListingType, request.Price)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create listing", err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"listing": listing,
	})
}
