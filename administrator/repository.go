package administrator

import "go-clean-arch-by-JR/models"

// Repository 123
type Repository interface {
	Get(hallID int, account string) (models.Administrator, error)
	GetListByHall(hallID int) ([]models.Administrator, error)
	Update(administratorData *models.Administrator)
}

// SidRepository 123
type SidRepository interface {
	StoreSid(sid string, administratorData string) error
	GetAdministratorDataBySid(sid string) (string, error)
	DeleteSid(sid string) error
}
