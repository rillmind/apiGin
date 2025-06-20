package product

type Service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return Service{
		repository: repo,
	}
}

func (ps *Service) GetProducts() ([]Product, error) {
	return ps.repository.GetProducts()
}

func (ps *Service) CreateProduct(product Product) (Product, error) {
	productID, err := ps.repository.CreateProduct(product)
	if err != nil {
		return Product{}, err
	}

	product.ID = productID

	return product, nil
}

func (ps *Service) GetProductByID(productID int) (*Product, error) {
	product, err := ps.repository.GetProductByID(productID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *Service) DeleteProductByID(productID int) (int64, error) {
	product, err := ps.repository.DeleteProductByID(productID)

	if err != nil {
		return 0, err
	}

	return product, err
}
