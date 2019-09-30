package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/japhethca/postit-api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func bindBodyOrJSON(ctx *gin.Context, v interface{}) error {
	if ctx.ContentType() == "application/json" {
		return ctx.ShouldBindJSON(v)
	}
	return ctx.ShouldBindWith(v, binding.FormPost)
}

// HashPassword hashes a password and returns error as nil if successful.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// SamePassword checks hashedPassword and password and returns true if both matches or false if not.
func samePassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}

func errorResponse(ctx interface{ JSON(int, interface{}) }, statuscode int, errDetail error, errMessage interface{}) {
	res := map[string]interface{}{
		"error":      errMessage,
		"status":     http.StatusText(statuscode),
		"statusCode": statuscode,
	}
	if os.Getenv("ENV") != "production" && errDetail != nil {
		res["detail"] = errDetail.Error()
	}
	ctx.JSON(statuscode, res)
}

func generateToken(user models.User) (string, error) {
	claim := tokenClaims{
		user.UserID,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "postit-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
