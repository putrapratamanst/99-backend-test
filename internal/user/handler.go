package user
import (
	"99-backend-exercise/internal/models"
	"99-backend-exercise/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
type Handler struct {
	userService Service
}
func NewHandler(userService Service) *Handler {
	return &Handler{
		userService: userService,
	}
}
func (h *Handler) GetUsers(c *gin.Context) {
	var request models.GetUsersRequest
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
func (h *Handler) CreateUser(c *gin.Context) {
	var request models.CreateUserRequest
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
