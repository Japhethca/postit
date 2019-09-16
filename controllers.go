package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

// Controller houses all handlers for routes
type Controller interface {
	// postit controllers
	GetPostit(*gin.Context)
	UpdatePostit(*gin.Context)
	GetAllPostit(*gin.Context)
	DeletePostit(*gin.Context)
	CreatePostit(*gin.Context)
}

type controllers struct {
	db *sql.DB
}

func (c *controllers) SetDB(db *sql.DB) {
	c.db = db
}

func (c *controllers) GetPostit(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{
		"message": "Awaiting implementation",
	})
}

func (c *controllers) UpdatePostit(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{
		"message": "Awaiting implementation",
	})
}

func (c *controllers) GetAllPostit(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{
		"message": "Awaiting implementation",
	})
}

func (c *controllers) CreatePostit(ctx *gin.Context) {
	var postit Postit
	err := bindBodyOrJSON(ctx, &postit)
	log.Println(err)
	log.Println(ctx.ContentType())
	ctx.JSON(http.StatusOK, postit)
}

func (c *controllers) DeletePostit(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{
		"message": "Awaiting implementation",
	})
}

func bindBodyOrJSON(ctx *gin.Context, v interface{}) error {
	if ctx.ContentType() == "application/json" {
		return ctx.ShouldBindJSON(v)
	}
	return ctx.ShouldBindWith(v, binding.FormPost)
}
