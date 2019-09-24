package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/db"
	"github.com/japhethca/postit-api/service/controllers"
	"net/http"
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
}

func (svc *APIService) Init() {
	svc.initControllers()
	svc.initRoutes()
}
