package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/yourname/user-service/internal/models"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrUserNotFound = errors.New("user not found")
	ErrEmailTaken   = errors.New("email already registered")
)

type UserRepository interface {
	Save(u *models.User) error
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name, email string) (*models.User, error) {
	if !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	}

	user := &models.User{
		ID:        generateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Save(user); err != nil {
		return nil, ErrEmailTaken
	}
	return user, nil
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func generateID() string {
	return fmt.Sprintf("usr_%d", time.Now().UnixNano())
}
