package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/internal/api/services"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func InitUserController(userService services.UserServiceInterface) *UserController {
	controller := &UserController{UserService: userService}
	return controller
}


func(c *UserController) UploadPhoto(ctx *gin.Context){
	fileHeader, err := ctx.FormFile("file")

	user, exists := ctx.Get("user")
    if !exists {
        ctx.AbortWithStatus(http.StatusUnauthorized)
        return
    }

	loggedUser := user.(*models.UserModel)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	

	if fileHeader == nil {
		sendError(ctx, 400, "file is required")
		return
	}


	resp,err  := c.UserService.UploadPhoto(loggedUser, fileHeader)

	if err != nil {
		sendError(ctx, 400, err.Error())
		return
	}

	sendSuccess(ctx, "success", resp)
	return
}