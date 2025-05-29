package main

import (
	"github.com/rillmind/apiGin/product"
	"github.com/rillmind/apiGin/user"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/db"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepository := user.NewUserRepository(dbConnection)
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	user.RegisterUserRoutes(server, &userController)

	productRepository := product.NewProductRepository(dbConnection)
	productService := product.NewProductService(productRepository)
	productController := product.NewProductController(productService)

	product.RegisterProductRoutes(server, &productController)

	server.Run(":2306")
}
