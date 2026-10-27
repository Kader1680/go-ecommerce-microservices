package main

import (
	"log"
	"net/http"
	"strings"
	"github.com/yourname/product-service/internal/handler"
	"github.com/yourname/product-service/internal/repository"
	"github.com/yourname/product-service/internal/service"
)

func main() {
	repo := repository.NewProductRepository()
	svc := service.NewProductService(repo)
	h := handler.NewProductHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.ListProducts(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/reserve") && r.Method == http.MethodPost {
			h.ReserveStock(w, r)
			return
		}
		if r.Method == http.MethodGet {
			h.GetProduct(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Product service running on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
