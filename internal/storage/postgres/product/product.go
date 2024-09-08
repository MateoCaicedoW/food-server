package product

import (
	"context"
	"database/sql"
	"fmt"
	"food-server/internal/entities/product"

	"github.com/gofrs/uuid/v5"
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

// GetByID returns a product by its id.
func (s *ProductRepository) GetByID(c context.Context, id uuid.UUID) (product.Single, error) {
	product := product.Single{}
	err := s.db.QueryRowContext(c, "SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return product, fmt.Errorf("failed to get product by id: %w", err)
	}

	return product, nil
}

// Update updates a product in the database.
func (s *ProductRepository) Update(c context.Context, p *product.Single) error {
	_, err := s.db.ExecContext(c, "UPDATE products SET name = $1, description = $2, price = $3, updated_at = now() WHERE id = $4", p.Name, p.Description, p.Price, p.ID)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}
