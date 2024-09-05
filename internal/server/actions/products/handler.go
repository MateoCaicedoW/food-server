package products

import (
	"food-server/internal/entities/product"
	"food-server/internal/json"
	"food-server/internal/service"
	"net/http"
)

type handler struct {
	productService service.ProductService
}

func (h handler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.List(r.Context())
	if err != nil {
		json.Response(w, json.Data{
			Content: json.Error{Message: err.Error()},
			Status:  http.StatusInternalServerError,
			OK:      false,
		})
		return
	}

	json.Response(w, json.Data{
		Content: products,
		Status:  http.StatusOK,
		OK:      true,
	})
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var product product.Single
	if err := json.Decode(r, &product); err != nil {
		json.Response(w, json.Data{
			Content: json.Error{Message: err.Error()},
			Status:  http.StatusBadRequest,
			OK:      false,
		})
		return
	}

	verrs, err := h.productService.ValidateAndSave(r.Context(), &product)
	if err != nil {
		json.Response(w, json.Data{
			Content: json.Error{Message: err.Error()},
			Status:  http.StatusInternalServerError,
			OK:      false,
		})

		return
	}

	if verrs.HasAny() {
		json.Response(w, json.Data{
			Content: verrs,
			Status:  http.StatusBadRequest,
			OK:      true,
		})

		return
	}

	json.Response(w, json.Data{
		Content: product,
		Status:  http.StatusCreated,
		OK:      true},
	)
}
