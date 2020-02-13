package main

import (
	"fmt"
	"go-clean-arch-by-JR/game/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	repo := repository.NewRepository(newMysql())
	result, err := repo.GetAgentToken("bbin-CNY", "041bef98cbe98074c8cd8f8b96f66b99")
	if err != nil {
		log.Printf("%#v", err)

	}
	// 建立mysql
	log.Printf("%#v", result)
}

func newMysql() *gorm.DB {
	connectName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "qwe123", "127.0.0.1", "3388", "brand")
	Conn, err := gorm.Open("mysql", connectName)
	Conn.LogMode(true)
	if err != nil {
		log.Fatalf("建立db連線失敗: %s", err.Error())
	}

	return Conn
}
