package usecase

import (
	"errors"
	"go-clean-arch-by-JR/game/mocks"
	"go-clean-arch-by-JR/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 測試範例，其他自己舉一反三
func Test_CheckToken_Have_Token(t *testing.T) {
	repo := new(mocks.Repository)
	repo.On("GetAgentToken", mock.Anything, mock.Anything).Return(models.Agent{UserName: "bbin-CNY", Token: "041bef98cbe98074c8cd8f8b96f66b99"}, nil)
	// 假設
	useCase := NewgameUsecase(repo)

	// 執行
	haveToken, err := useCase.CheckToken("bbin-CNY", "041bef98cbe98074c8cd8f8b96f66b99")

	// 斷言
	assert.Equal(t, true, haveToken)
	assert.Equal(t, nil, err)
}

func Test_CheckToken_Agent_Null(t *testing.T) {
	repo := new(mocks.Repository)
	repo.On("GetAgentToken", mock.Anything, mock.Anything).Return(models.Agent{}, nil)

	// 假設
	useCase := NewgameUsecase(repo)

	// 執行
	haveToken, err := useCase.CheckToken("bbin-CNY", "041bef98cbe98074c8cd8f8b96f66b99")

	// 斷言
	assert.Equal(t, false, haveToken)
	assert.Equal(t, nil, err)
}

func Test_CheckToken_Error(t *testing.T) {
	repo := new(mocks.Repository)
	repo.On("GetAgentToken", mock.Anything, mock.Anything).Return(models.Agent{}, errors.New("test"))

	// 假設
	useCase := NewgameUsecase(repo)

	// 執行
	haveToken, err := useCase.CheckToken("bbin-CNY", "041bef98cbe98074c8cd8f8b96f66b99")

	// 斷言
	assert.Equal(t, false, haveToken)
	assert.Error(t, err)
}
