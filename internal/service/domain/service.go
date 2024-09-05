package domain

import (
	"food-server/internal/service"
	"food-server/internal/service/domain/product"
	"food-server/internal/storage"
)

func NewService(r *storage.Repository) *service.Service {
	return &service.Service{
		Product: product.NewProductService(r.Product),
	}
}
