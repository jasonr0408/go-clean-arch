package repository

import (
	"go-clean-arch-by-JR/administrator"

	"github.com/gomodule/redigo/redis"
	pkgErrors "github.com/pkg/errors"
)

type redisAdministratorRepository struct {
	Conn *redis.Pool
}

func NewRedisAdministratorRepository(Conn *redis.Pool) administrator.SidRepository {
	return &redisAdministratorRepository{Conn}
}

func (n *redisAdministratorRepository) StoreSid(sid string, administratorData string) (err error) {
	conn := n.Conn.Get()
	// errcheck ignore
	defer conn.Close()

	_, err = conn.Do("SET", "sid:"+sid, administratorData)
	if err != nil {
		return pkgErrors.Wrap(err, "redis儲存sid失敗")
	}

	_, err = conn.Do("EXPIRE", sid, 30000)
	if err != nil {
		return pkgErrors.Wrap(err, "redis設定生存時間失敗")
	}

	return
}

func (n *redisAdministratorRepository) GetAdministratorDataBySid(sid string) (administratorData string, err error) {
	conn := n.Conn.Get()
	defer conn.Close()

	redisReturn, err := conn.Do("GET", "sid:"+sid)
	if err != nil {
		return "", pkgErrors.Wrap(err, "redis取administrator失敗1")
	}

	if redisReturn != nil {
		administratorData, err = redis.String(redisReturn, err)
		if err != nil {
			return "", pkgErrors.Wrap(err, "redis取administrator失敗2")
		}
	}

	return
}

func (n *redisAdministratorRepository) DeleteSid(sid string) (err error) {
	conn := n.Conn.Get()
	defer conn.Close()

	_, err = conn.Do("DEL", "sid:"+sid)
	if err != nil {
		return pkgErrors.Wrap(err, "redis刪除sid失敗")
	}

	return
}
