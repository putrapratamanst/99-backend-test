package user

import (
	"99-backend-exercise/internal/models"
	"99-backend-exercise/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler defines the HTTP handler for user endpoints
type Handler struct {
	userService Service
}

// NewHandler creates a new user handler
func NewHandler(userService Service) *Handler {
	return &Handler{
		userService: userService,
	}
}

// GetUsers handles GET /users
func (h *Handler) GetUsers(c *gin.Context) {
	var request models.GetUsersRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&request); err != nil {
		utils.RespondWithValidationError(c, err)
		return
	}

	users, err := h.userService.GetUsers(request)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get users", err)
		return
	}

	response := map[string]interface{}{
		"users": users,
	}
	utils.RespondWithSuccess(c, response)
}

// GetUserByID handles GET /users/{id}
func (h *Handler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found", err)
		return
	}

	response := map[string]interface{}{
		"user": user,
	}
	utils.RespondWithSuccess(c, response)
}

// CreateUser handles POST /users
func (h *Handler) CreateUser(c *gin.Context) {
	var request models.CreateUserRequest

	// Bind request body (form data)
	if err := c.ShouldBind(&request); err != nil {
		fmt.Println(err)

		utils.RespondWithValidationError(c, err)
		return
	}

	user, err := h.userService.CreateUser(request)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	response := map[string]interface{}{
		"user": user,
	}
	utils.RespondWithSuccess(c, response)
}
