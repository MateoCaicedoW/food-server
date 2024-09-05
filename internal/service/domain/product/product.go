package product

import (
	"context"
	"food-server/internal/entities/product"
	"food-server/internal/storage"
	"food-server/system/validate"
)

type ProductService struct {
	ProductRepository storage.ProductRepository
}

func NewProductService(r storage.ProductRepository) ProductService {
	return ProductService{ProductRepository: r}
}

func (s ProductService) List(c context.Context) (product.All, error) {
	return s.ProductRepository.List(
		c,
	)
}

func (s ProductService) Validate(p product.Single) validate.Errors {
	errors := validate.Errors{}
	errors.Required("name", p.Name)
	errors.Required("price", p.Price)
	errors.GreatherThan("price", p.Price, 0)
	errors.Required("description", p.Description)
	return errors
}

func (s ProductService) ValidateAndSave(c context.Context, p *product.Single) (validate.Errors, error) {
	errs := s.Validate(*p)
	if errs.HasAny() {
		return errs, nil
	}

	err := s.ProductRepository.Create(c, p)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
