package storage

import (
	"context"
	"food-server/internal/entities/product"
)

type ProductRepository interface {
	List(context.Context) (product.All, error)
	Create(context.Context, *product.Single) error
}

type Repository struct {
	Product ProductRepository
}
