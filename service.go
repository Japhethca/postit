package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type service struct {
	router      *gin.Engine
	db          *sql.DB
	controllers Controller
}

func (svc *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	svc.router.ServeHTTP(w, r)
}

func (svc *service) InitRoutes() {
	apiV1 := svc.router.Group("/api/v1")
	apiV1.GET("/postit/:id", svc.controllers.GetPostit)
	apiV1.GET("/postit/", svc.controllers.GetAllPostit)
	apiV1.POST("/postit/", svc.controllers.CreatePostit)
	apiV1.PUT("/postit/:id", svc.controllers.UpdatePostit)
	apiV1.DELETE("/postit:id/", svc.controllers.DeletePostit)
}
