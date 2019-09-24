package models

import (
	"time"
)

// Postit model
type Postit struct {
	ID        int       `json:"id" form:"id"`
	Content   string    `json:"content" form:"content"`
	UserID    int       `json:"userId" form:"userId"`
	GroupID   int       `json:"groups" form:"groups"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}
