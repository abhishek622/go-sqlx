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
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, "INSERT INTO products (name, image, category, description,rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)", p)

		if err != nil {
			return fmt.Errorf("Error while adding product: %w", err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("Error getting last insert ID: %w", err)
		}

		p.ID = id
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating product: %w", err)
	}

	return p, nil
}

func (ms *MySQLStorer) GetProduct(ctx context.Context, id int64) (*Product, error) {
	var p Product
	err := ms.db.GetContext(ctx, &p, "SELECT * FROM products where id=?", id)
	if err != nil {
		return nil, fmt.Errorf("Error getting product: %w", err)
	}

	return &p, nil
}

func (ms *MySQLStorer) ListProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	err := ms.db.SelectContext(ctx, &products, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("Error listing products: %w", err)
	}

	return products, nil
}

func (ms *MySQLStorer) UpdateProduct(ctx context.Context, p *Product) (*Product, error) {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, "UPDATE products SET name=:name, image=:image, category=:category, description=:description, rating=:rating, num_reviews=:num_reviews, price=:price, count_in_stock=:count_in_stock, updated_at=:updated_at WHERE id=:id", p)
		if err != nil {
			return fmt.Errorf("error updating product: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error while updating product: %w", err)
	}

	return p, nil
}

func (ms *MySQLStorer) DeleteProduct(ctx context.Context, id int64) error {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, "DELETE FROM products WHERE id=?", id)
		if err != nil {
			return fmt.Errorf("Error deleting product items: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("Error while deleting product: %w", err)
	}

	return nil
}

func (ms *MySQLStorer) CreateOrder(ctx context.Context, o *Order) (*Order, error) {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		// insert into order
		order, err := createOrder(ctx, tx, o)
		if err != nil {
			return fmt.Errorf("Error creating order: %w", err)
		}

		for _, oi := range o.Items {
			oi.OrderID = order.ID
			err = createOrderItem(ctx, tx, oi)
			if err != nil {
				return fmt.Errorf("Error creating order item: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating order: %w", err)
	}

	return o, nil
}

func createOrder(ctx context.Context, tx *sqlx.Tx, o *Order) (*Order, error) {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO orders (payment_method, tax_price, shipping_price, total_price) VALUES (:payment_method, :tax_price, :shipping_price, :total_price)", o)
	if err != nil {
		return nil, fmt.Errorf("Error inserting order: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Error getting last order ID: %w", err)
	}

	o.ID = id

	return o, nil
}

func createOrderItem(ctx context.Context, tx *sqlx.Tx, oi OrderItem) error {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO order_items (order_id, product_id, name, quantity, image, price) VALUES (:order_id, :product_id, :name, :quantity, :image, :price)", oi)
	if err != nil {
		return fmt.Errorf("Error inserting order item: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error getting last order item ID: %w", err)
	}

	oi.ID = id

	return nil
}

func (ms *MySQLStorer) GetOrder(ctx context.Context, id int64) (*Order, error) {
	var o Order
	err := ms.db.GetContext(ctx, &o, "SELECT * FROM orders WHERE id=?", id)
	if err != nil {
		return nil, fmt.Errorf("Error getting order: %w", err)
	}

	var items []OrderItem
	err = ms.db.SelectContext(ctx, &items, "SELECT * FROM order_items WHERE order_id=?", id)
	if err != nil {
		return nil, fmt.Errorf("Error getting order items: %w", err)
	}

	o.Items = items
	return &o, nil
}

func (ms *MySQLStorer) ListOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	err := ms.db.SelectContext(ctx, &orders, "SELECT * FROM orders")
	if err != nil {
		return nil, fmt.Errorf("Error listing orders: %w", err)
	}

	for i := range orders {
		var items []OrderItem
		err = ms.db.SelectContext(ctx, &items, "SELECT * FROM order_items WHERE order_id=?", orders[i].ID)
		if err != nil {
			return nil, fmt.Errorf("Error getting order items: %w", err)
		}
		orders[i].Items = items
	}

	return orders, nil
}

func (ms *MySQLStorer) DeleteOrder(ctx context.Context, id int64) error {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, "DELETE FROM order_items WHERE order_id=?", id)
		if err != nil {
			return fmt.Errorf("Error deleting order items: %w", err)
		}
		_, err = tx.ExecContext(ctx, "DELETE FROM order WHERE id=?", id)
		if err != nil {
			return fmt.Errorf("Error deleting order: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("Error while deleting order: %w", err)
	}

	return nil
}

// user
func (ms *MySQLStorer) CreateUser(ctx context.Context, u *User) (*User, error) {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, "INSERT INTO products (name, email, password, is_admin) VALUES (:name, :email, :password, :is_admin)", u)

		if err != nil {
			return fmt.Errorf("Error while adding product: %w", err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("Error getting last insert ID: %w", err)
		}

		u.ID = id
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating user: %w", err)
	}

	return u, nil
}

func (ms *MySQLStorer) GetUser(ctx context.Context, email string) (*User, error) {
	var u User
	err := ms.db.GetContext(ctx, &u, "SELECT * FROM users where email=?", email)
	if err != nil {
		return nil, fmt.Errorf("Error getting user: %w", err)
	}

	return &u, nil
}
func (ms *MySQLStorer) GetUserById(ctx context.Context, id int64) (*User, error) {
	var u User
	err := ms.db.GetContext(ctx, &u, "SELECT * FROM users where id=?", id)
	if err != nil {
		return nil, fmt.Errorf("Error getting user: %w", err)
	}

	return &u, nil
}

func (ms *MySQLStorer) ListUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := ms.db.SelectContext(ctx, &users, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("Error listing users: %w", err)
	}

	return users, nil
}

func (ms *MySQLStorer) UpdateUser(ctx context.Context, u *User) (*User, error) {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, "UPDATE users SET name=:name, email=:email, password=:password, is_admin=:is_admin, updated_at=:updated_at WHERE id=:id", u)
		if err != nil {
			return fmt.Errorf("error updating user: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error while updating user: %w", err)
	}

	return u, nil
}

func (ms *MySQLStorer) DeleteUser(ctx context.Context, id int64) error {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, "DELETE FROM users WHERE id=?", id)
		if err != nil {
			return fmt.Errorf("Error deleting user items: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("Error while deleting user: %w", err)
	}

	return nil
}

// helper function
func (ms MySQLStorer) execTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := ms.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Error starting transaction: %w", err)
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Error rolling back transaction: %w", rbErr)
		}
		return fmt.Errorf("Error in transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Error committing transaction: %w", err)
	}

	return nil
}
