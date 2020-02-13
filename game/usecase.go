package game

// Usecase
type Usecase interface {
	CheckToken(agentName, token string) (bool, error)
}
