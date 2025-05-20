package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/controller"
	"github.com/rillmind/apiGin/db"
	"github.com/rillmind/apiGin/repository"
	"github.com/rillmind/apiGin/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)

	productUsecase := usecase.NewProductUsecase(productRepository)

	productController := controller.NewProductController(productUsecase)

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Run(":2306")
}
