package mocks

import (
	"go-clean-arch-by-JR/models"

	"github.com/stretchr/testify/mock"
)

// Repository 123
type Repository struct {
	mock.Mock
}

func (_m *Repository) Get(hallID int, account string) (models.Administrator, error) {

	return models.Administrator{}, nil
}

func (_m *Repository) GetListByHall(hallID int) ([]models.Administrator, error) {

	return []models.Administrator{}, nil
}

func (_m *Repository) Update(administratorData *models.Administrator) {
}

// SidRepository 123
type SidRepository struct {
	mock.Mock
}

func (_m *SidRepository) StoreSid(sid string, administratorData string) error {
	return nil
}

func (_m *SidRepository) GetAdministratorDataBySid(sid string) (string, error) {
	return `{"ID": 123}`, nil
}

func (_m *SidRepository) DeleteSid(sid string) error {
	return nil
}
