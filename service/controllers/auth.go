package controllers

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/db"
	"github.com/japhethca/postit-api/db/dao"
	"github.com/japhethca/postit-api/models"
	"github.com/japhethca/postit-api/service/validator"
	"github.com/lib/pq"
	"net/http"
	"os"
)

// Auth holds all controllers for authentication routes
type Auth struct {
	db *sql.DB
}

// SetDB satisfies db.Setter interface for assigning db to object
func (c *Auth) SetDB(db *sql.DB) {
	c.db = db
}

// Signup handles user registration
func (c *Auth) Signup(ctx *gin.Context) {
	var signupFields validator.SignupFields
	bindBodyOrJSON(ctx, &signupFields)
	verrs := validator.SignupValidator(&signupFields)
	if len(verrs) > 0 {
		ErrorResponse(ctx, http.StatusBadRequest, nil, verrs)
		return
	}

	if !sameString(signupFields.Password, signupFields.ConfirmPassword) {
		ErrorResponse(ctx, http.StatusBadRequest, nil, "Passwords did not match")
		return
	}

	userDAO := dao.UserDAO{DB: c.db}
	user := models.User{
		Email:     signupFields.Email,
		Firstname: signupFields.FirstName,
		Lastname:  signupFields.LastName,
	}
	hashedPassword, _ := hashPassword(signupFields.Password)
	user.Password = hashedPassword

	newUser, err := userDAO.CreateUser(user)
	if err != nil {
		dbError, ok := err.(*pq.Error)
		if ok && dbError.Code == db.DatabaseUniqueViolation {
			ErrorResponse(ctx, http.StatusConflict, nil, "User with this credentials already exists")
			return
		}
		ErrorResponse(ctx, http.StatusInternalServerError, err, "Its not you, its us. We are working hard to resolve this.")
		return
	}
	token, _ := generateToken(newUser)
	ctx.JSON(http.StatusCreated, gin.H{
		"user":  newUser,
		"token": token,
	})
}

// Login handles user login
func (c *Auth) Login(ctx *gin.Context) {
	var loginFields validator.LoginFields
	bindBodyOrJSON(ctx, &loginFields)
	verrs := validator.LoginValidator(&loginFields)
	if len(verrs) > 0 {
		ErrorResponse(ctx, http.StatusBadRequest, nil, verrs)
		return
	}
	userManager := dao.UserDAO{DB: c.db}
	user, err := userManager.GetUserByEmail(loginFields.Email)
	if err != nil || !samePassword(user.Password, loginFields.Password) {
		ErrorResponse(ctx, http.StatusUnauthorized, err, "Invalid login credentials")
		return
	}
	token, _ := generateToken(user)
	ctx.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.UserID,
			"email": user.Email,
		},
		"token": token,
	})
}

func (c *Auth) TestController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to Postit API, have fun creating your stuffs.",
	})
}

type TokenClaims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

func generateToken(user models.User) (string, error) {
	claim := TokenClaims{
		user.UserID,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "postit-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
