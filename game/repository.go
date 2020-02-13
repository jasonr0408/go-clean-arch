package game

import "go-clean-arch-by-JR/models"

// Repository 123
type Repository interface {
	GetAgentToken(agentName, token string) (game models.Agent, err error)
}
