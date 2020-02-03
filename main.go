package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	administrator "go-clean-arch-by-JR/administrator/delivery/http"
	administratorUseRepo "go-clean-arch-by-JR/administrator/repository"
	administratorUse "go-clean-arch-by-JR/administrator/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

func main() {
	// 載入router root
	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9487")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// cors
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	sqlConn := newMysql()
	redisConn := newRedis()

	// administrator api
	administratorRepository := administratorUseRepo.NewMysqlAdministratorRepository(sqlConn)
	administratorSidRepository := administratorUseRepo.NewRedisAdministratorRepository(redisConn)
	administratorUsecase := administratorUse.NewAdministratorUsecase(administratorRepository, administratorSidRepository)
	administrator.NewAdministratorHandler(r, administratorUsecase)

	return r
}

// 建立mysql
func newMysql() *gorm.DB {
	connectName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "qwe123", "127.0.0.1", "3388", "collection")
	Conn, err := gorm.Open("mysql", connectName)
	Conn.LogMode(true)
	if err != nil {
		log.Fatalf("建立db連線失敗: %s", err.Error())
	}

	return Conn
}

// 建立redis
func newRedis() *redis.Pool {
	Conn := &redis.Pool{
		Wait:        true,
		MaxIdle:     20,
		MaxActive:   2000,
		IdleTimeout: 10 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6378")
			if err != nil {
				return nil, err
			}
			// _, err = c.Do("AUTH", "123")
			// if err != nil {
			// 	return nil, err
			// }

			c.Do("SELECT", 3)

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return Conn
}
