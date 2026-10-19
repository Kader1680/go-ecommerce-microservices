package main

import (
	"log"
	"net/http"

	"github.com/yourname/order-service/internal/handler"
	"github.com/yourname/order-service/internal/repository"
	"github.com/yourname/order-service/internal/service"
)

func main() {
	repo := repository.NewOrderRepository()
	svc := service.NewOrderService(repo)
	h := handler.NewOrderHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateOrder(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	mux.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetOrder(w, r)
			return
		}
		if r.Method == http.MethodPost {
			h.CancelOrder(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("🚀 Order service running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
