package product

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{
		repository: repo,
	}
}

func (ps *ProductService) GetProducts() ([]Product, error) {
	return ps.repository.GetProducts()
}

func (ps *ProductService) CreateProduct(product Product) (Product, error) {
	productID, err := ps.repository.CreateProduct(product)
	if err != nil {
		return Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (ps *ProductService) GetProductByID(productID int) (*Product, error) {
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
