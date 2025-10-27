package repository

import (
	"database/sql"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	// Read from existing table `product` (singular) and map column `product_name` to `name`
	query := "SELECT id, product_name AS name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return nil, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			return nil, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}
