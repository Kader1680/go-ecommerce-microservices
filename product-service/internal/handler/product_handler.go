package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/yourname/product-service/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, _ := h.service.ListProducts()
	respondJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")
	p, err := h.service.GetProduct(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "product not found")
		return
	}
	respondJSON(w, http.StatusOK, p)
}

type reserveStockRequest struct {
	Quantity int `json:"quantity"`
}

// Called by other services (e.g. Order Service) to reserve stock - THIS is inter-service communication.
func (h *ProductHandler) ReserveStock(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")
	id = strings.TrimSuffix(id, "/reserve")

	var req reserveStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}

	if err := h.service.ReserveStock(id, req.Quantity); err != nil {
		if errors.Is(err, service.ErrProductNotFound) {
			respondError(w, http.StatusNotFound, err.Error())
			return
		}
		respondError(w, http.StatusConflict, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"status": "reserved"})
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
