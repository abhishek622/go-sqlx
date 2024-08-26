package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MySQLStorer struct {
	db *sqlx.DB
}

func NewMySQLStorer(db *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{
		db: db,
	}
}

func (ms *MySQLStorer) CreateProduct(ctx context.Context, p *Product) (*Product, error) {
	res, err := ms.db.NamedExecContext(ctx, "INSERT INTO products (name, image, category, description,rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)", p)

	if err != nil {
		return nil, fmt.Errorf("Error while adding product: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Error getting last insert ID: %w", err)
	}

	p.ID = id
	return p, nil
}

func (ms *MySQLStorer) GetProduct(ctx context.Context, id int64) (*Product, error) {
	var p Product
	err := ms.db.GetContext(ctx, &p, "SELECT * FROM products where id=?")
	if err != nil {
		return nil, fmt.Errorf("Error getting product: %w", err)
	}

	return &p, nil
}

func (ms *MySQLStorer) ListProducts(ctx context.Context) ([]*Product, error) {
	var products []*Product
	err := ms.db.SelectContext(ctx, &products, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("Error listing products: %w", err)
	}

	return products, nil
}

func (ms *MySQLStorer) updateProduct(ctx context.Context, p *Product) (*Product, error) {
	_, err := ms.db.NamedExecContext(ctx, "UPDATE products SET name = :name, image = :image, category = :category, description = :description, rating = :rating, num_reviews = :num_reviews, price = :price, count_in_stock = :count_in_stock, updated_at = now() WHERE id = :id", p)
	if err != nil {
		return nil, fmt.Errorf("Error updating product: %w", err)
	}

	return p, nil
}

func (ms *MySQLStorer) deleteProduct(ctx context.Context, id int64) error {
	_, err := ms.db.ExecContext(ctx, "DELETE FROM products WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("Error while deleting product: %w", err)
	}

	return nil
}
