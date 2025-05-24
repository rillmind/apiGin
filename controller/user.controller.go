package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/service"
)

type userController struct {
	service.UserService
}

func NewUserController(service service.UserService) userController {
	return userController{
		UserService: service,
	}
}

func (uc *userController) GetUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}
