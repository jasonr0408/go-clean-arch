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
	return models.Agent{UserName: "bbin-CNY", Token: "041bef98cbe98074c8cd8f8b96f66b99"}, nil
}
