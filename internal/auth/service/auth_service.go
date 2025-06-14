package service

import (
	"github.com/zaidalmaghfur/go-app/internal/auth/domain"
	"github.com/zaidalmaghfur/go-app/internal/auth/repository/pgsql"
	"github.com/zaidalmaghfur/go-app/pkg"
)

type AuthService interface {
	Register(name, email, password string) (*domain.User, error)
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo pgsql.UserRepository
}

func NewAuthService(userRepo pgsql.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(name, email, password string) (*domain.User, error) {
	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// Jangan return password!
	user.Password = ""

	return user, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !pkg.CheckPasswordHash(password, user.Password) {
		return "", pkg.ErrInvalidCredentials
	}

	token, err := pkg.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
