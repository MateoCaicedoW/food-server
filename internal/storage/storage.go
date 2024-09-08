package storage

import (
	"context"
	"food-server/internal/entities/product"

	"github.com/gofrs/uuid/v5"
)

type ProductRepository interface {
	List(context.Context) (product.All, error)
	Create(context.Context, *product.Single) error
	GetByID(context.Context, uuid.UUID) (product.Single, error)
	Update(context.Context, *product.Single) error
}

type Repository struct {
	Product ProductRepository
}
