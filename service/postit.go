package service

import (
	"github.com/gin-gonic/gin"
	"github.com/japhethca/postit-api/db/dao"
	"github.com/japhethca/postit-api/models"
	"github.com/japhethca/postit-api/service/validator"
	"net/http"
	"time"
)

func (svc *Service) createPostit(ctx *gin.Context) {
	var postitFields validator.CreatPostitFields
	bindBodyOrJSON(ctx, &postitFields)
	if verr := validator.CreatePostitValidator(&postitFields); len(verr) > 0 {
		errorResponse(ctx, http.StatusBadRequest, nil, verr)
		return
	}
	postit := models.Postit{
		Content:   postitFields.Content,
		Color:     postitFields.Color,
		UserID:    postitFields.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newPostit, err := dao.NewPostit(svc.db, postit)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err, "Its not you, its us. We are working hard to resolve this.")
		return
	}
	ctx.JSON(http.StatusCreated, newPostit)
}
