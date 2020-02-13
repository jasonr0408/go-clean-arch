package repository

import (
	"errors"
	"go-clean-arch-by-JR/game"
	"go-clean-arch-by-JR/models"

	"github.com/jinzhu/gorm"
)

type repository struct {
	Conn *gorm.DB
}

func NewRepository(Conn *gorm.DB) game.Repository {
	return &repository{
		Conn,
	}
}

func (m *repository) GetAgentToken(agentName, token string) (game models.Agent, err error) {
	if err := m.Conn.Where("UserName = ? AND Token = ?", agentName, token).Find(&game).Error; err != nil {
		return game, errors.New("取不到Agent資料")
	}

	return game, nil
}
