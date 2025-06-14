package service

import (
	"github.com/zaidalmaghfur/go-app/internal/master/agama/domain"
	"github.com/zaidalmaghfur/go-app/internal/master/agama/repository/pgsql"
)

type AgamaService interface {
	Create(name string) (*domain.Agama, error)
	GetByID(id uint) (*domain.Agama, error)
	Update(id uint, name string) (*domain.Agama, error)
	Delete(id uint) error
	GetAll() ([]domain.Agama, error)
	GetAllPaginated(offset, limit int) ([]domain.Agama, int, error)
}

type agamaService struct {
	agamaRepo pgsql.AgamaRepository
}

func NewAgamaService(repo pgsql.AgamaRepository) AgamaService {
	return &agamaService{repo}
}

func (s *agamaService) Create(name string) (*domain.Agama, error) {
	agama := &domain.Agama{Name: name}
	err := s.agamaRepo.Create(agama)
	return agama, err
}

func (s *agamaService) GetByID(id uint) (*domain.Agama, error) {
	return s.agamaRepo.FindByID(id)
}

func (s *agamaService) Update(id uint, name string) (*domain.Agama, error) {
	agama, err := s.agamaRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	agama.Name = name
	err = s.agamaRepo.Update(agama)
	return agama, err
}

func (s *agamaService) Delete(id uint) error {
	return s.agamaRepo.Delete(id)
}

func (s *agamaService) GetAll() ([]domain.Agama, error) {
	return s.agamaRepo.FindAll()
}

func (s *agamaService) GetAllPaginated(offset, limit int) ([]domain.Agama, int, error) {
	return s.agamaRepo.FindAllPaginated(offset, limit)
}
