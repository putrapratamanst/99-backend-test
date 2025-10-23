package user
import (
	"99-backend-exercise/internal/models"
	"gorm.io/gorm"
)
type Repository interface {
	GetAll(offset, limit int) ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
	Count() (int64, error)
}
type repository struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
func (r *repository) GetAll(offset, limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}
func (r *repository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
func (r *repository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return count, err
}
