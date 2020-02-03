package usecase

import (
	"encoding/json"
	"errors"
	"go-clean-arch-by-JR/administrator"
	"go-clean-arch-by-JR/models"

	pkgErrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type administratorUsecase struct {
	administratorRepo administrator.Repository
	sidRepo           administrator.SidRepository
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewAdministratorUsecase(a administrator.Repository, s administrator.SidRepository) administrator.Usecase {
	return &administratorUsecase{
		administratorRepo: a,
		sidRepo:           s,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

// func checkPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func (a *administratorUsecase) Login(hallID int, account, password string) (sid string, err error) {
	// get user data
	administratorData, err := a.administratorRepo.Get(hallID, account)
	if err != nil {
		return
	}

	// check password && empty account
	if administratorData.Account == "" {
		return "", errors.New("沒有這個使用者")
	}
	if administratorData.Password != password {
		return "", errors.New("密碼錯誤")
	}

	// encodeing sid
	sid, err = hashPassword(string(administratorData.HallID) + administratorData.Account + administratorData.Password)
	if err != nil {
		return "", pkgErrors.Wrap(err, "sid加密錯誤")
	}

	// json encode administratorData
	jsonAdministratorData, err := json.Marshal(administratorData)

	// store redis
	err = a.sidRepo.StoreSid(sid, string(jsonAdministratorData))
	if err != nil {
		return
	}
	// return sid

	return
}

func (a *administratorUsecase) GetInfo(sid string) (administratorData *models.Administrator, err error) {
	// 用sid取redis的資料
	jsonAdministratorData, err := a.sidRepo.GetAdministratorDataBySid(sid)
	if err != nil {
		return
	}
	if jsonAdministratorData == "" {
		return administratorData, pkgErrors.Wrap(err, "請重新登入")
	}

	err = json.Unmarshal([]byte(jsonAdministratorData), &administratorData)
	if err != nil {
		return administratorData, pkgErrors.Wrap(err, "json轉換格式失敗")
	}

	if administratorData.ID == 0 {
		return administratorData, pkgErrors.Wrap(errors.New("沒有sid"), "json轉換格式失敗")
	}

	return
}

func (a *administratorUsecase) Logout(sid string) (err error) {
	err = a.sidRepo.DeleteSid(sid)

	return
}

func (a *administratorUsecase) GetList(hallID int) (administratorList []models.Administrator, err error) {
	administratorList, err = a.administratorRepo.GetListByHall(hallID)

	return
}

func (a *administratorUsecase) Edit(sid string, administratorData *models.Administrator) (err error) {
	// 用sid取redis的資料
	operator, err := a.GetInfo(sid)
	if err != nil {
		return
	}
	// 確認是否可以修改
	if (operator.HallID != administratorData.HallID) && (operator.SalesID != "A") {
		return pkgErrors.Wrap(err, "沒有權限")
	}

	// 修改
	a.administratorRepo.Update(administratorData)

	return
}
