package http

import (
	"go-clean-arch-by-JR/administrator"
	"go-clean-arch-by-JR/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v9"
)

type AdministratorHandler struct {
	AdministratorUsecase administrator.Usecase
}

type login struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewAdministratorHandler(r *gin.Engine, a administrator.Usecase) {
	handler := &AdministratorHandler{a}

	r.POST("/administrator/login", handler.Login)
	r.GET("/administrator/info", handler.GetInfo)
	r.POST("/administrator/logout", handler.Logout)
	r.GET("/administrator/list", handler.GetList)
	r.PUT("/administrator", handler.Edit)
}

func isRequestValid(m *login) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *AdministratorHandler) Login(c *gin.Context) {
	var login login
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
		return
	}

	if ok, err := isRequestValid(&login); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
		return
	}

	sid, err := a.AdministratorUsecase.Login(1, login.Account, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			// "code": 20000,
			"data": map[string]string{"sid": sid},
		})
	}
}

func (a *AdministratorHandler) GetSidData(c *gin.Context) (sid string, administratorData *models.Administrator, errMsg string) {
	// 拿到sid
	sid = c.GetHeader("X-Token")
	if sid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": "sid錯誤"})
		return "", administratorData, "sid錯誤"
	}
	// call usecase
	administratorData, err := a.AdministratorUsecase.GetInfo(sid)
	if err != nil {
		return "", administratorData, err.Error()
	}

	return
}

func (a *AdministratorHandler) GetInfo(c *gin.Context) {
	// call usecase
	_, administratorData, errMsg := a.GetSidData(c)
	if errMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": errMsg})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			// "code": 20000,
			"data": map[string]interface{}{
				"hallID":  administratorData.HallID,
				"salesID": administratorData.SalesID,
				"account": administratorData.Account,
			},
		})
	}
}

func (a *AdministratorHandler) Logout(c *gin.Context) {
	sid, _, errMsg := a.GetSidData(c)
	if errMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": errMsg})
		return
	}
	// call usecase
	err := a.AdministratorUsecase.Logout(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			// "code": 20000,
			"data": "success",
		})
	}
}

func (a *AdministratorHandler) GetList(c *gin.Context) {
	_, _, errMsg := a.GetSidData(c)
	if errMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": errMsg})
		return
	}

	// 拿到hallID
	sHallID := c.Query("hallID")
	if sHallID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": "請重新登入!"})
		return
	}
	hallID, _ := strconv.Atoi(sHallID)
	// call usecase
	administratorList, err := a.AdministratorUsecase.GetList(hallID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": administratorList,
		})
	}
}

func (a *AdministratorHandler) Edit(c *gin.Context) {
	sid, _, errMsg := a.GetSidData(c)
	if errMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": errMsg})
		return
	}
	var data models.Administrator
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = a.AdministratorUsecase.Edit(sid, &data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	}
}
