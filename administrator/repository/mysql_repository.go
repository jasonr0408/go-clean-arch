package repository

import (
	"go-clean-arch-by-JR/administrator"
	"go-clean-arch-by-JR/models"
	"fmt"

	"github.com/jinzhu/gorm"
	pkgErrors "github.com/pkg/errors"
)

type mysqlAdministratorRepository struct {
	Conn *gorm.DB
}

func NewMysqlAdministratorRepository(Conn *gorm.DB) administrator.Repository {
	return &mysqlAdministratorRepository{Conn}
}

func (m *mysqlAdministratorRepository) Get(hallID int, account string) (administrator models.Administrator, err error) {
	if err := m.Conn.Where("hallID = ? AND account = ?", hallID, account).Find(&administrator).Error; err != nil {
		return administrator, pkgErrors.Wrap(err, "取mysql的Administrator錯誤")
	}

	return administrator, nil
}

func (m *mysqlAdministratorRepository) GetListByHall(hallID int) (administratorList []models.Administrator, err error) {
	sql := fmt.Sprintf("SELECT * FROM `administrator_list` WHERE `hallID` = '%d'", hallID)

	rows, err := m.Conn.Raw(sql).Rows()

	// rows, err := m.Conn.Table("administrator_list").Where("hallID = ?", hallID).Rows()
	// if err != nil {
	// 	return nil, pkgErrors.Wrap(err, "取mysql的Administrator列表錯誤1")
	// }
	// defer rows.Close()

	for rows.Next() {
		var administrator models.Administrator
		if err := rows.Scan(&administrator.ID, &administrator.HallID, &administrator.SalesID, &administrator.Account, &administrator.Password, &administrator.Name); err != nil {
			return nil, pkgErrors.Wrap(err, "取mysql的Administrator列表錯誤2")
		}
		administratorList = append(administratorList, administrator)
	}

	return
}

func (m *mysqlAdministratorRepository) Update(administratorData *models.Administrator) {
	m.Conn.Save(&administratorData)
}
