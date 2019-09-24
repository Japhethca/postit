package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func bindBodyOrJSON(ctx *gin.Context, v interface{}) error {
	if ctx.ContentType() == "application/json" {
		return ctx.ShouldBindJSON(v)
	}
	return ctx.ShouldBindWith(v, binding.FormPost)
}

func sameString(first, second string) bool {
	if first == second {
		return true
	}
	return false
}

// HashPassword hashes a password and returns error as nil if successful.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// SamePassword checks hashedPassword and password and returns true if both matches or false if not.
func SamePassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}

type jsonResponse interface {
	JSON(int, interface{})
}

func errorResponse(ctx jsonResponse, statuscode int, err interface{}) {
	ctx.JSON(statuscode, map[string]interface{}{
		"error":      err,
		"status":     http.StatusText(statuscode),
		"statusCode": statuscode,
	})
}
