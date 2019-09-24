package validator

import (
	"github.com/japhethca/postit-api/db/models"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

// SignupValidator validates signup routes
func SignupValidator(user *models.User) url.Values {
	rules := govalidator.MapData{
		"firstName": []string{"required"},
		"lastName":  []string{"required"},
		"email":     []string{"required", "email"},
	}
	return user.Validate(rules)
}
