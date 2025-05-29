package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(server *gin.Engine, pc *ProductController) {
	server.GET("/products", pc.GetProducts)
	server.GET("/product/:productID", pc.GetProductByID)
	server.POST("/product", pc.CreateProduct)
	server.DELETE("/product/:productID", pc.DeleteProductByID)
}
