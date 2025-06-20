package main

import (
	"github.com/rillmind/apiGin/auth"
	"github.com/rillmind/apiGin/product"
	"github.com/rillmind/apiGin/user"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/db"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	authRepository := auth.NewRepository(dbConnection)
	authServise := auth.NewService(authRepository)
	authController := auth.NewController(authServise)

	auth.RegisterRoutes(server, &authController)

	userRepository := user.NewRepository(dbConnection)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	user.RegisterRoutes(server, &userController)

	productRepository := product.NewRepository(dbConnection)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	product.RegisterRoutes(server, &productController)

	server.Run(":2306")
}
