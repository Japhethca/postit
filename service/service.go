package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Service is an API service
type Service struct {
	router *gin.Engine
	db     *sql.DB
}

// New creates new api Service
func New(router *gin.Engine, db *sql.DB) *Service {
	return &Service{router: router, db: db}
}

func (svc *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svc.router.ServeHTTP(w, r)
}

func (svc *Service) apiRoutes() {
	apiV1 := svc.router.Group("/api/v1")
	apiV1.GET("/", svc.withAuth(svc.testController))
	apiV1.POST("/auth/signup", svc.signup)
	apiV1.POST("/auth/login", svc.login)
	apiV1.POST("/postit/", svc.createPostit)
}

func (svc *Service) Init() {
	svc.apiRoutes()
}

func (svc *Service) testController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to Postit API, have fun creating your stuffs.",
	})
}
