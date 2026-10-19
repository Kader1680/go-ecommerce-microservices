package repository

import (
	"errors"
	"sync"

	"github.com/yourname/user-service/internal/models"
)

var (
	ErrNotFound      = errors.New("user not found")
	ErrEmailTaken    = errors.New("email already registered")
)

type UserRepository struct {
	mu    sync.RWMutex
	users map[string]*models.User
	// secondary index so we don't scan the whole map on every signup
	byEmail map[string]string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:   make(map[string]*models.User),
		byEmail: make(map[string]string),
	}
}

func (r *UserRepository) Save(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if existingID, ok := r.byEmail[u.Email]; ok && existingID != u.ID {
		return ErrEmailTaken
	}
	r.users[u.ID] = u
	r.byEmail[u.Email] = u.ID
	return nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.byEmail[email]
	if !ok {
		return nil, ErrNotFound
	}
	return r.users[id], nil
}
