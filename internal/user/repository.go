package user

import (
	"99-backend-exercise/internal/models"

	"gorm.io/gorm"
)

// Repository defines the interface for user data operations
type Repository interface {
	GetAll(offset, limit int) ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
	Count() (int64, error)
}

// repository is the concrete implementation of Repository interface
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new user repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// GetAll retrieves users with pagination, sorted by creation date (descending)
func (r *repository) GetAll(offset, limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// GetByID retrieves a user by ID
func (r *repository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create creates a new user
func (r *repository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Count returns the total number of users
func (r *repository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return count, err
}
