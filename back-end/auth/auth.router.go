package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine, ac *Controller) {
	server.POST("/login", ac.Login)
}
