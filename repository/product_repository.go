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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id=$1")
	if err != nil {
		return nil, err
	}

	var product model.Product
	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = query.Exec(id)
	return err
}

func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	res, err := pr.connection.Exec("UPDATE product SET product_name=$1, price=$2 WHERE id=$3", product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
