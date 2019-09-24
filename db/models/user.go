package models

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

// User model
type User struct {
	UserID    int    `json:"userId" form:"userId"`
	Firstname string `json:"firstName" form:"firstName"`
	Lastname  string `json:"lastName" form:"lastName"`
	Email     string `json:"email" form:"email"`
}

// Validate validates user model given a rule
func (user *User) Validate(rules govalidator.MapData) url.Values {
	ops := govalidator.Options{
		Data:  user,
		Rules: rules,
	}
	return govalidator.New(ops).ValidateStruct()
}
