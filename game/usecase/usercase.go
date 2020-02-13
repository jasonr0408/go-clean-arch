package usecase

import (
	"go-clean-arch-by-JR/game"
)

type gameUsecase struct {
	Repo game.Repository
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewgameUsecase(a game.Repository) game.Usecase {
	return &gameUsecase{
		Repo: a,
	}
}

func (_a *gameUsecase) CheckToken(agentName, token string) (bool, error) {
	// 撈db get agent token
	agent, err := _a.Repo.GetAgentToken(agentName, token)
	if err != nil {
		return false, err
	}

	// 比對
	if agent.Token == "" {
		return false, nil
	}

	return true, nil
}
