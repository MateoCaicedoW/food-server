package products

import (
	"food-server/internal/service"

	"github.com/leapkit/leapkit/core/server"
)

func RegisterRoutes(productService service.ProductService, router server.Router) {
	handler := handler{productService}
	router.HandleFunc("GET /list", handler.List)
	router.HandleFunc("POST /create", handler.Create)
}
