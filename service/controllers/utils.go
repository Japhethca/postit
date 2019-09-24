package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func bindBodyOrJSON(ctx *gin.Context, v interface{}) error {
	if ctx.ContentType() == "application/json" {
		return ctx.ShouldBindJSON(v)
	}
	return ctx.ShouldBindWith(v, binding.FormPost)
}
