package mocks

import (
	"go-clean-arch-by-JR/models"

	"github.com/stretchr/testify/mock"
)

// Usecase
type Usecase struct {
	mock.Mock
}

func (_m *Usecase) Login(hallID int, account, password string) (string, error) {
	return "", nil
}

func (_m *Usecase) GetInfo(sid string) (*models.Administrator, error) {
	return &models.Administrator{}, nil
}

func (_m *Usecase) Logout(sid string) error {
	return nil
}

func (_m *Usecase) GetList(hallID int) ([]models.Administrator, error) {
	return []models.Administrator{}, nil
}

func (_m *Usecase) Edit(sid string, administratorData *models.Administrator) error {
	return nil
}
