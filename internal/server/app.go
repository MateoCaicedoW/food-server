package server

import (
	"cmp"
	"food-server/internal/config"
	"food-server/internal/server/actions/products"
	"food-server/internal/server/middleware"
	"food-server/internal/service/domain"
	"food-server/internal/storage/postgres"

	"net/http"

	"os"

	"github.com/leapkit/leapkit/core/server"
)

// Server interface exposes the methods
// needed to start the server in the cmd/app package
type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() Server {
	// Creating a new server instance with the
	// default host and port values.
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3001")),
	)

	r.Use(middleware.CORS)

	db, err := config.DB()
	if err != nil {
		panic(err)
	}

	repositories := postgres.NewRepository(db)
	services := domain.NewService(repositories)

	r.Group("/api", func(r server.Router) {
		// Registering the routes for the products
		// action group.
		r.Group("/products", func(r server.Router) {
			products.RegisterRoutes(services.Product, r)
		})

	})

	return r

}
