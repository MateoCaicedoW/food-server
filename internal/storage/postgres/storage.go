package postgres

import (
	"database/sql"
	"food-server/internal/storage"
	"food-server/internal/storage/postgres/product"
)

func NewRepository(db *sql.DB) *storage.Repository {
	return &storage.Repository{
		Product: product.NewProductRepository(db),
	}
}
