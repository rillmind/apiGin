package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, pc *Controller) {
	server.GET("/products", pc.GetProducts)
	server.GET("/product/:productID", pc.GetProductByID)
	server.POST("/product", pc.CreateProduct)
	server.DELETE("/product/:productID", pc.DeleteProductByID)
}
