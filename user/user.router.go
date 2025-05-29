package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(server *gin.Engine, uc *UserController) {
	server.GET("/users", uc.GetUsers)
	server.GET("/user/:userID", uc.GetUserByID)
	server.POST("/user", uc.CreateUser)
	server.DELETE("/user/:userID", uc.DeleteUserByID)
}
