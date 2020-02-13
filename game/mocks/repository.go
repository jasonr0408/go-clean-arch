package mocks

import (
	"go-clean-arch-by-JR/models"

	"github.com/stretchr/testify/mock"
)

// Repository 123
type Repository struct {
	mock.Mock
}

func (_m *Repository) GetAgentToken(agentName, token string) (game models.Agent, err error) {
	args := _m.Called(agentName, token)
	return args.Get(0).(models.Agent), args.Error(1)

	// return models.Agent{UserName: "bbin-CNY", Token: "041bef98cbe98074c8cd8f8b96f66b99"}, nil
}
