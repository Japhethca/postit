package models

// User model
type User struct {
	UserID    int    `json:"userId" form:"userId"`
	Firstname string `json:"firstName" form:"firstName"`
	Lastname  string `json:"lastName" form:"lastName"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}
