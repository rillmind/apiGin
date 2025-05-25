package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/controller"
)

func RegisterUserRoutes(server *gin.Engine, uc *controller.UserController) {
	server.GET("/users", uc.GetUsers)
	server.POST("/user", uc.CreateUser)
	server.GET("/user/:userID", uc.GetUserByID)
}
