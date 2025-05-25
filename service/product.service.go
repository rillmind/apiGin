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

func (ps *ProductService) GetProducts() ([]model.Product, error) {
	return ps.repository.GetProducts()
}

func (ps *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := ps.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (ps *ProductService) GetProductByID(productID int) (*model.Product, error) {
	product, err := ps.repository.GetProductByID(productID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) DeleteProductByID(productID int) (int64, error) {
	product, err := ps.repository.DeleteProductByID(productID)

	if err != nil {
		return 0, err
	}

	return product, err
}
