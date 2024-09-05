package config

import (
	"cmp"
	"os"

	"github.com/leapkit/leapkit/core/db"
)

var (

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(
		cmp.Or(os.Getenv("DATABASE_URL"), "postgres://postgres:postgres@127.0.0.1:5432/food_server?sslmode=disable"),
	)
)
