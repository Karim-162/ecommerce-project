package cmd

import (
	"ecommerceProject47/globalRouter"
	"ecommerceProject47/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.Getproducts))
	mux.Handle("POST /createproduct", http.HandlerFunc(handlers.CreateProduct))

	globalrouter := globalRouter.GlobalRouter(mux)
	fmt.Println("server running on :8080")
	err := http.ListenAndServe(":8080", globalrouter)
	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}
