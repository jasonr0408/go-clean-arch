package http

import (
	"errors"
	"go-clean-arch-by-JR/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameUsecase game.Usecase
}

type login struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewGameHandler(r *gin.Engine, a game.Usecase) {
	handler := &GameHandler{a}

	r.GET("/gamelink", handler.GetGameLink)
}

func (_a *GameHandler) checkToken(c *gin.Context) error {
	// 先取get參數
	// agent table的UserName、Token
	agentName := c.Query("agentname")
	if agentName == "" {
		// 寫log
		return errors.New("")
	}

	token := c.Query("token")
	if token == "" {
		// 寫log
		return errors.New("")
	}
	haveToken, err := _a.GameUsecase.CheckToken(agentName, token)
	if err != nil {
		return err
	}

	if !haveToken {
		return errors.New("沒有這個Agent")
	}

	return nil
}

func (_a *GameHandler) GetGameLink(c *gin.Context) {
	err := _a.checkToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": 111, "error_message": "Token驗證失敗", "execution_time": 0, "data": ""})
	}

	// 先取get參數
	gameID := c.Query("gameid")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": 111, "error_message": "少了gameid參數", "execution_time": 0, "data": ""})
		return
	}

	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": 111, "error_message": "少了gameid參數", "execution_time": 0, "data": ""})
		return
	}

	// var lang string
	// lang = c.Query("lang")

	// var login login
	// err := c.BindJSON(&login)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	// 	return
	// }

	// if ok, err := isRequestValid(&login); !ok {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	// 	return
	// }

	// sid, err := a.GameUsecase.Login(1, login.Account, login.Password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": true, "errMsg": err.Error()})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		// "code": 20000,
	// 		"data": map[string]string{"sid": sid},
	// 	})
	// }
}
