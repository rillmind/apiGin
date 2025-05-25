package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/model"
	"github.com/rillmind/apiGin/service"
)

type UserController struct {
	service.UserService
}

func NewUserController(service service.UserService) UserController {
	return UserController{
		UserService: service,
	}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
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

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userID")

	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.UserService.GetUserByID(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		response := model.Response{
			Message: "Usuário não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
