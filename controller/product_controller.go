package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/model"
	"github.com/rillmind/apiGin/usecase"
)

type productController struct {
	usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		ProductUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")

	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUsecase.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")

	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUsecase.DeleteProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == 0 {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
