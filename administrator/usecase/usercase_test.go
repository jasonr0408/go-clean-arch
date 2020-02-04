package usecase

import (
	"testing"

	administratorUseRepo "go-clean-arch-by-JR/administrator/mocks"
	"go-clean-arch-by-JR/models"

	"github.com/stretchr/testify/assert"
)

// 測試範例，其他自己舉一反三
func Test_GetInfo(t *testing.T) {
	// 1a
	administratorRepository := administratorUseRepo.Repository{}
	administratorSidRepository := administratorUseRepo.SidRepository{}
	administratorUsecase := NewAdministratorUsecase(&administratorRepository, &administratorSidRepository)

	// 2a
	result, err := administratorUsecase.GetInfo("123")

	// 3a
	assert.Equal(t, &models.Administrator{ID: 123}, result)
	assert.Equal(t, nil, err)
}
