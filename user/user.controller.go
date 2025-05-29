package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/response"
)

type UserController struct {
	UserService
}

func NewUserController(service UserService) UserController {
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
	var user User

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

	userResp := UserResponse{
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
		response := response.New{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		response := response.New{
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
		response := response.New{
			Message: "Usuário não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("userID")

	if id == "" {
		response := response.New{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		response := response.New{
			Message: "ID precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.UserService.DeleteUserByID(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == 0 {
		response := response.New{
			Message: "Usuário não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
