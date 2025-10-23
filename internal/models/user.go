package models

type User struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name" binding:"required"`
	Timestamp
}
type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: ToMicroseconds(u.CreatedAt),
		UpdatedAt: ToMicroseconds(u.UpdatedAt),
	}
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}
type GetUsersRequest struct {
	PaginationRequest
}
