package service

import (
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

type tokenClaims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

func (svc *Service) withAuth(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		if authToken == "" {
			errorResponse(ctx, http.StatusUnauthorized, nil, "Invalid authorization token")
			return
		}
		claims := tokenClaims{}
		token, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if token.Valid {
			userDao := dao.UserDAO{DB: svc.db}
			if _, err := userDao.GetUserByID(claims.UserID); err != nil {
				errorResponse(ctx, http.StatusUnauthorized, err, "You are not authorized to access this route.")
				return
			}
			handler(ctx)
			return
		}
		if err != nil {
			errorResponse(ctx, http.StatusUnauthorized, err, "Invalid authorization token")
		}
	}
}

// Signup handles user registration
func (svc *Service) signup(ctx *gin.Context) {
	var signupFields validator.SignupFields
	bindBodyOrJSON(ctx, &signupFields)
	verrs := validator.SignupValidator(&signupFields)
	if len(verrs) > 0 {
		errorResponse(ctx, http.StatusBadRequest, nil, verrs)
		return
	}

	if signupFields.Password != signupFields.ConfirmPassword {
		errorResponse(ctx, http.StatusBadRequest, nil, "Passwords did not match")
		return
	}

	userDAO := dao.UserDAO{DB: svc.db}
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
			errorResponse(ctx, http.StatusConflict, nil, "User with this credentials already exists")
			return
		}
		errorResponse(ctx, http.StatusInternalServerError, err, "Its not you, its us. We are working hard to resolve this.")
		return
	}
	token, _ := generateToken(newUser)
	ctx.JSON(http.StatusCreated, gin.H{
		"user":  newUser,
		"token": token,
	})
}

// Login handles user login
func (svc *Service) login(ctx *gin.Context) {
	var loginFields validator.LoginFields
	bindBodyOrJSON(ctx, &loginFields)
	verrs := validator.LoginValidator(&loginFields)
	if len(verrs) > 0 {
		errorResponse(ctx, http.StatusBadRequest, nil, verrs)
		return
	}
	userManager := dao.UserDAO{DB: svc.db}
	user, err := userManager.GetUserByEmail(loginFields.Email)
	if err != nil || !samePassword(user.Password, loginFields.Password) {
		errorResponse(ctx, http.StatusUnauthorized, err, "Invalid login credentials")
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
