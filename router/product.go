package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/controller"
)

func RegisterProductRoutes(server *gin.Engine, pc *controller.ProductController) {
	server.GET("/products", pc.GetProducts)
	server.GET("/product/:productID", pc.GetProductByID)
	server.POST("/product", pc.CreateProduct)
	server.DELETE("/product/:productID", pc.DeleteProductByID)
}
