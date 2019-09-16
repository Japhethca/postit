package main

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

// PostitGroup model
type PostitGroup struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

// User model
type User struct {
	UserID    int    `json:"userId" form:"userId"`
	Firstname string `json:"firstName" form:"firstName"`
	Lastname  string `json:"lastName" form:"lastName"`
	Email     string `json:"email" form:"email"`
}
