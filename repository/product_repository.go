package repository

import (
	"database/sql"
	"fmt"

	"github.com/rillmind/apiGin/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	var productList []model.Product
	var productObj model.Product

	query := "select id, product_name, price from product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Print(err)
		return []model.Product{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}
