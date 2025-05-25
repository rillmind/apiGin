package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/model"
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

func (uc *userController) CreateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := uc.UserService.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	userResp := model.UserResponse{
		ID:       insertedUser.ID,
		Name:     insertedUser.Name,
		Username: insertedUser.Username,
		Email:    insertedUser.Email,
	}

	ctx.JSON(http.StatusCreated, userResp)
}
