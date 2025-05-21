package usecase

import (
	"github.com/rillmind/apiGin/model"
	"github.com/rillmind/apiGin/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (pu *ProductUsecase) GetProductByID(productID int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(productID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductByID(productID int) (int64, error) {
	product, err := pu.repository.DeleteProductByID(productID)

	if err != nil {
		return 0, err
	}

	return product, err
}
