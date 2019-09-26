package service

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/db"
	"github.com/japhethca/postit-api/db/dao"
	"github.com/japhethca/postit-api/service/controllers"
	"net/http"
	"os"
)

type APIService struct {
	router *gin.Engine
	db     *sql.DB
}

var authController controllers.Auth

// New creates new api service
func New(router *gin.Engine, db *sql.DB) *APIService {
	return &APIService{
		router: router,
		db:     db,
	}
}

func (svc *APIService) initControllers() {
	svc.setControllerDB(&authController)

}

func (svc *APIService) setControllerDB(controller db.Setter) {
	controller.SetDB(svc.db)
}

func (svc *APIService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svc.router.ServeHTTP(w, r)
}

func (svc *APIService) initRoutes() {
	apiV1 := svc.router.Group("/api/v1")
	apiV1.POST("/auth/signup", authController.Signup)
	apiV1.POST("/auth/login", authController.Login)
	apiV1.GET("test", svc.withAuth(authController.TestController))
}

func (svc *APIService) Init() {
	svc.initControllers()
	svc.initRoutes()
}

func (svc *APIService) withAuth(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		if authToken == "" {
			controllers.ErrorResponse(ctx, http.StatusUnauthorized, nil, "Invalid authorization token")
			return
		}
		claims := controllers.TokenClaims{}
		token, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if token.Valid {
			userDao := dao.UserDAO{DB: svc.db}
			if _, err := userDao.GetUserByID(claims.UserID); err != nil {
				controllers.ErrorResponse(ctx, http.StatusUnauthorized, err, "You are not authorized to access this route.")
				return
			}
			handler(ctx)
			return
		}
		if err != nil {
			controllers.ErrorResponse(ctx, http.StatusUnauthorized, err, "Invalid authorization token")
		}
	}
}
