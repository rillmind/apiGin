package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rillmind/apiGin/response"
)

type Controller struct {
	Service
}

func NewController(service Service) Controller {
	return Controller{
		ProductService: service,
	}
}

func (pc *Controller) GetProducts(ctx *gin.Context) {
	products, err := pc.ProductService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *Controller) CreateProduct(ctx *gin.Context) {
	var product Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := pc.ProductService.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (pc *Controller) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")

	if id == "" {
		response := response.New{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)

	if err != nil {
		response := response.New{
			Message: "ID precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.ProductService.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := response.New{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *Controller) DeleteProductByID(ctx *gin.Context) {
	id := ctx.Param("productID")

	if id == "" {
		response := response.New{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)

	if err != nil {
		response := response.New{
			Message: "ID precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.ProductService.DeleteProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == 0 {
		response := response.New{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
