package validator

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type SignupFields struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginFields struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupValidator validates signup route
func SignupValidator(fields *SignupFields) url.Values {
	rules := govalidator.MapData{
		"firstName":       []string{"required"},
		"lastName":        []string{"required"},
		"email":           []string{"required", "email"},
		"password":        []string{"required", "min:6"},
		"confirmPassword": []string{"required"},
	}
	return validate(fields, rules)
}

// LoginValidator validates signin route
func LoginValidator(fields *LoginFields) url.Values {
	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}
	return validate(fields, rules)
}

func validate(data interface{}, rules govalidator.MapData) url.Values {
	ops := govalidator.Options{
		Data:  data,
		Rules: rules,
	}
	return govalidator.New(ops).ValidateStruct()
}
