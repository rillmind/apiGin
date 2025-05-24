package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/controller"
	"github.com/rillmind/apiGin/db"
	"github.com/rillmind/apiGin/repository"
	"github.com/rillmind/apiGin/service"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(dbConnection)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	server.GET("/users", userController.GetUsers)

	productRepository := repository.NewProductRepository(dbConnection)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productID", productController.GetProductByID)
	server.POST("/product", productController.CreateProduct)
	server.DELETE("/product/:productID", productController.DeleteProductByID)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Run(":2306")
}
