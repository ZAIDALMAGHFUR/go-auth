package mysql

import (
	"github.com/zaidalmaghfur/go-app/config"
	"github.com/zaidalmaghfur/go-app/internal/auth/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *domain.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
