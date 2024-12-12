package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/api/services"
)

type AuthController struct {
	AuthService services.AuthServiceInterface
}

func InitAuthController(authService services.AuthServiceInterface) *AuthController {
	controller := &AuthController{AuthService: authService}
	return controller
}

func (c *AuthController) SignIn(ctx *gin.Context){

	dto := dtos.SignInDto{}

    err := ctx.BindJSON(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    err = dto.Validate()
    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    data, err := c.AuthService.SignIn(&dto)

    if err != nil {
       sendError(ctx, 401, err.Error())
        return 
    }

    sendSuccess(ctx, "success", data)
    return
}


func (c *AuthController) SignUp(ctx *gin.Context){

	dto := dtos.CreateUserDto{}

    err := ctx.BindJSON(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    err = dto.Validate()
    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    data, err := c.AuthService.SignUp(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

   sendSuccess(ctx, "success", data)
    return
}