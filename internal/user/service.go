package user
import (
	"99-backend-exercise/internal/models"
	"errors"
)
type Service interface {
	GetUsers(request models.GetUsersRequest) ([]models.UserResponse, error)
	GetUserByID(id int) (*models.UserResponse, error)
	CreateUser(request models.CreateUserRequest) (*models.UserResponse, error)
}
type service struct {
	userRepo Repository
}
func NewService(userRepo Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}
func (s *service) GetUsers(request models.GetUsersRequest) ([]models.UserResponse, error) {
	offset := request.GetOffset()
	limit := request.GetPageSize()
	users, err := s.userRepo.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}
	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}
	return responses, nil
}
func (s *service) GetUserByID(id int) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	response := user.ToResponse()
	return &response, nil
}
func (s *service) CreateUser(request models.CreateUserRequest) (*models.UserResponse, error) {
	user := &models.User{
		Name: request.Name,
	}
	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	response := user.ToResponse()
	return &response, nil
}
