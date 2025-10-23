package models

import (
	"time"
)

type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}
type PaginationRequest struct {
	PageNum  int `form:"page_num" json:"page_num" binding:"min=1"`
	PageSize int `form:"page_size" json:"page_size" binding:"min=1,max=100"`
}

func (p *PaginationRequest) GetPageNum() int {
	if p.PageNum <= 0 {
		return 1
	}
	return p.PageNum
}
func (p *PaginationRequest) GetPageSize() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}
func (p *PaginationRequest) GetOffset() int {
	return (p.GetPageNum() - 1) * p.GetPageSize()
}

type Timestamp struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func ToMicroseconds(t time.Time) int64 {
	return t.UnixNano() / 1000
}
