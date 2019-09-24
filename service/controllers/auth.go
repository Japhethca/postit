package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/db"
	"github.com/japhethca/postit-api/db/manager"
	"github.com/japhethca/postit-api/db/models"
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

// Signup controller signs up user
func (c *Auth) Signup(ctx *gin.Context) {
	var user models.User
	bindBodyOrJSON(ctx, &user)
	verrs := validator.SignupValidator(&user)
	if len(verrs) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      verrs,
			"status":     http.StatusText(http.StatusBadRequest),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	userManager := manager.UserManager{DB: c.db}
	newUser, err := userManager.CreateUser(user)
	if err != nil {
		dbError := err.(*pq.Error)
		if dbError.Code == db.DatabaseUniqueViolation {
			ctx.JSON(http.StatusConflict, gin.H{
				"error":      "User with this credentials already exists",
				"status":     http.StatusText(http.StatusConflict),
				"statusCode": http.StatusConflict,
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":      "Something unexpected happened. Its not you, its us.",
			"status":     http.StatusText(http.StatusInternalServerError),
			"statusCode": http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}
