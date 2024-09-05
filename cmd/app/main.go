package main

import (
	"fmt"
	"food-server/internal/server"
	"net/http"

	// Load environment variables

	_ "github.com/leapkit/leapkit/core/tools/envload"

	// Load the database driver
	_ "github.com/lib/pq"
)

func main() {
	s := server.New()
	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}
