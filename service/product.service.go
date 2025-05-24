package service

import (
	"github.com/rillmind/apiGin/model"
	"github.com/rillmind/apiGin/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return ProductService{
		repository: repo,
	}
}

func (pu *ProductService) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (pu *ProductService) GetProductByID(productID int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(productID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductService) DeleteProductByID(productID int) (int64, error) {
	product, err := pu.repository.DeleteProductByID(productID)

	if err != nil {
		return 0, err
	}

	return product, err
}
