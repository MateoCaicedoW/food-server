package service

import (
	"context"
	"food-server/internal/entities/product"
	"food-server/system/validate"

	"github.com/gofrs/uuid/v5"
)

type ProductService interface {
	List(context.Context) (product.All, error)
	ValidateAndSave(context.Context, *product.Single) (validate.Errors, error)
	GetByID(context.Context, uuid.UUID) (product.Single, error)
	ValidateAndUpdate(context.Context, *product.Single) (validate.Errors, error)
}

type Service struct {
	Product ProductService
}
