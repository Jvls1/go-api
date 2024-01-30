package database

import (
	"database/sql"

	"github.com/jvls1/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func newProduct(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT id, name, price, category_id FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductDB) CreateCategory(category *entity.Product) (string, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
