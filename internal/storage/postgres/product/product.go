package product

import (
	"context"
	"database/sql"
	"fmt"
	"food-server/internal/entities/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// List returns all products from the database.
func (s *ProductRepository) List(c context.Context) (product.All, error) {
	rows, err := s.db.QueryContext(c, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	defer rows.Close()

	products := product.All{}
	for rows.Next() {
		product := product.Single{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}
	return products, nil
}

// Create creates a new product in the database
// and put the id in the product.Single struct.
func (s *ProductRepository) Create(c context.Context, p *product.Single) error {
	err := s.db.QueryRowContext(c, "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at", p.Name, p.Description, p.Price).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}
