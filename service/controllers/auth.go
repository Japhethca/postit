package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/db"
	"github.com/japhethca/postit-api/db/dao"
	"github.com/japhethca/postit-api/models"
	"github.com/japhethca/postit-api/service/validator"
	"github.com/lib/pq"
	"net/http"
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
		errorResponse(ctx, http.StatusBadRequest, verrs)
		return
	}

	if !sameString(signupFields.Password, signupFields.ConfirmPassword) {
		errorResponse(ctx, http.StatusBadRequest, "Passwords did not match")
		return
	}

	userDAO := dao.UserDAO{DB: c.db}
	user := models.User{
		Email:     signupFields.Email,
		Firstname: signupFields.FirstName,
		Lastname:  signupFields.LastName,
	}
	hashedPassword, _ := HashPassword(signupFields.Password)
	user.Password = hashedPassword

	newUser, err := userDAO.CreateUser(user)
	if err != nil {
		dbError, ok := err.(*pq.Error)
		if ok && dbError.Code == db.DatabaseUniqueViolation {
			errorResponse(ctx, http.StatusConflict, "User with this credentials already exists")
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, "Its not you, its us. We are working hard to resolve this.")
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}

// Login handles user login
func (c *Auth) Login(ctx *gin.Context) {
	var loginFields validator.LoginFields
	bindBodyOrJSON(ctx, &loginFields)
	verrs := validator.LoginValidator(&loginFields)
	if len(verrs) > 0 {
		errorResponse(ctx, http.StatusBadRequest, verrs)
		return
	}
	userManager := dao.UserDAO{DB: c.db}
	user, err := userManager.GetUserByEmail(loginFields.Email)
	if err != nil || !SamePassword(user.Password, loginFields.Password) {
		errorResponse(ctx, http.StatusUnauthorized, "Invalid login credentials")
		return
	}
	ctx.JSON(http.StatusOK, user)
}
