package service

import (
	"context"
	"food-server/internal/entities/product"
	"food-server/system/validate"
)

type ProductService interface {
	List(context.Context) (product.All, error)
	ValidateAndSave(context.Context, *product.Single) (validate.Errors, error)
}

type Service struct {
	Product ProductService
}
