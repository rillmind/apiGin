package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine, uc *Controller) {
	server.GET("/users", uc.GetUsers)
	server.GET("/user/:userID", uc.GetUserByID)
	server.POST("/user", uc.CreateUser)
	server.DELETE("/user/:userID", uc.DeleteUserByID)
}
