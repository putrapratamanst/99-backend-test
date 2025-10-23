package models

// User represents the user entity as defined in the API specification
type User struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name" binding:"required"`
	Timestamp
}

// UserResponse represents the user response format with microsecond timestamps
type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"` // microseconds
	UpdatedAt int64  `json:"updated_at"` // microseconds
}

// ToResponse converts User to UserResponse with microsecond timestamps
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: ToMicroseconds(u.CreatedAt),
		UpdatedAt: ToMicroseconds(u.UpdatedAt),
	}
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

// GetUsersRequest represents the request parameters for getting users
type GetUsersRequest struct {
	PaginationRequest
}
