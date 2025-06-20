package product

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	connection *sql.DB
}

func NewRepository(connection *sql.DB) Repository {
	return Repository{
		connection: connection,
	}
}

func (pr *Repository) GetProducts() ([]Product, error) {
	var productList []Product
	var productObj Product

	query := "select id, product_name, price from product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Print(err)
		return []Product{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *Repository) CreateProduct(product Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare(
		"insert into product" +
			"(product_name, price)" +
			"values ($1, $2) returning id",
	)

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *Repository) GetProductByID(productID int) (*Product, error) {
	var product Product

	query, err := pr.connection.Prepare("select * from product where id = $1")

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	err = query.QueryRow(productID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *Repository) DeleteProductByID(productID int) (int64, error) {
	query, err := pr.connection.Prepare("delete from product where id = $1")

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	defer query.Close()

	result, err := query.Exec(productID)

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	return rowsAffected, nil
}
