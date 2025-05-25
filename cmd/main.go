package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/controller"
	"github.com/rillmind/apiGin/db"
	"github.com/rillmind/apiGin/repository"
	"github.com/rillmind/apiGin/router"
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

	router.RegisterUserRoutes(server, &userController)

	productRepository := repository.NewProductRepository(dbConnection)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router.RegisterProductRoutes(server, &productController)

	server.Run(":2306")
}
