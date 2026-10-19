package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/yourname/user-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

type registerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}

	user, err := h.service.Register(req.Name, req.Email)
	if err != nil {
		if errors.Is(err, service.ErrInvalidEmail) {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, service.ErrEmailTaken) {
			respondError(w, http.StatusConflict, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, "internal error")
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	user, err := h.service.GetUser(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "user not found")
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}
