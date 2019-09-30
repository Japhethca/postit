package models

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

// Postit model
type Postit struct {
	ID         int       `json:"id" form:"id"`
	Content    string    `json:"content" form:"content"`
	Color      string    `json:"color" form:"color"`
	UserID     int       `json:"userId" form:"userId"`
	CategoryID null.Int  `json:"category" form:"category"`
	CreatedAt  time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" form:"updatedAt"`
}
