package pgsql

import (
	"github.com/zaidalmaghfur/go-app/config"
	"github.com/zaidalmaghfur/go-app/internal/master/agama/domain"
)

type AgamaRepository interface {
	Create(agama *domain.Agama) error
	FindByID(id uint) (*domain.Agama, error)
	Update(agama *domain.Agama) error
	Delete(id uint) error
	FindAll() ([]domain.Agama, error)
}

type agamaRepository struct{}

func NewAgamaRepository() AgamaRepository {
	return &agamaRepository{}
}

func (r *agamaRepository) Create(agama *domain.Agama) error {
	return config.DB.Create(agama).Error
}

func (r *agamaRepository) FindByID(id uint) (*domain.Agama, error) {
	var agama domain.Agama
	err := config.DB.First(&agama, id).Error
	return &agama, err
}

func (r *agamaRepository) Update(agama *domain.Agama) error {
	return config.DB.Save(agama).Error
}

func (r *agamaRepository) Delete(id uint) error {
	return config.DB.Delete(&domain.Agama{}, id).Error
}

func (r *agamaRepository) FindAll() ([]domain.Agama, error) {
	var agama []domain.Agama
	err := config.DB.Find(&agama).Error
	return agama, err
}