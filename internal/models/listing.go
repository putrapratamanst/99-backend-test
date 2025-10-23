package models

type Listing struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int    `gorm:"not null;index" json:"user_id" binding:"required"`
	Price       int    `gorm:"not null" json:"price" binding:"required,min=1"`
	ListingType string `gorm:"not null" json:"listing_type" binding:"required,oneof=rent sale"`
	Timestamp
}
type ListingResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (l *Listing) ToResponse() ListingResponse {
	return ListingResponse{
		ID:          l.ID,
		UserID:      l.UserID,
		ListingType: l.ListingType,
		Price:       l.Price,
		CreatedAt:   ToMicroseconds(l.CreatedAt),
		UpdatedAt:   ToMicroseconds(l.UpdatedAt),
	}
}

type CreateListingRequest struct {
	UserID      int    `json:"user_id" form:"user_id" binding:"required"`
	ListingType string `json:"listing_type" form:"listing_type" binding:"required,oneof=rent sale"`
	Price       int    `json:"price" form:"price" binding:"required,min=1"`
}
type GetListingsRequest struct {
	PaginationRequest
	UserID *int `form:"user_id" json:"user_id,omitempty"`
}
type PublicListingResponse struct {
	ID          int          `json:"id"`
	ListingType string       `json:"listing_type"`
	Price       int          `json:"price"`
	CreatedAt   int64        `json:"created_at"`
	UpdatedAt   int64        `json:"updated_at"`
	User        UserResponse `json:"user"`
}

func (l *Listing) ToPublicResponse(user User) PublicListingResponse {
	return PublicListingResponse{
		ID:          l.ID,
		ListingType: l.ListingType,
		Price:       l.Price,
		CreatedAt:   ToMicroseconds(l.CreatedAt),
		UpdatedAt:   ToMicroseconds(l.UpdatedAt),
		User:        user.ToResponse(),
	}
}
