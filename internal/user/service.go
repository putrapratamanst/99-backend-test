package user

import (
	"99-backend-exercise/internal/models"
	"errors"
)

// Service defines the interface for user business logic
type Service interface {
	GetUsers(request models.GetUsersRequest) ([]models.UserResponse, error)
	GetUserByID(id int) (*models.UserResponse, error)
	CreateUser(request models.CreateUserRequest) (*models.UserResponse, error)
}

// service is the concrete implementation of Service interface
type service struct {
	userRepo Repository
}

// NewService creates a new user service
func NewService(userRepo Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}

// GetUsers retrieves users with pagination
func (s *service) GetUsers(request models.GetUsersRequest) ([]models.UserResponse, error) {
	offset := request.GetOffset()
	limit := request.GetPageSize()

	users, err := s.userRepo.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	// Convert to response format
	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	return responses, nil
}

// GetUserByID retrieves a user by ID
func (s *service) GetUserByID(id int) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	response := user.ToResponse()
	return &response, nil
}

// CreateUser creates a new user
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
