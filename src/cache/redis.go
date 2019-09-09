package cache

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var gRedis redis.Conn = nil

func InitRedis() error {
	fmt.Println("Init Redis...")
	var err error = nil
	gRedis, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	fmt.Println("Init Redis...ok")
	return nil
}

func AddKeyValue(key string, val string) error {
	if gRedis != nil {
		gRedis.Do("set", key, val)
	}
	return nil
}

func GetValue(key string) (string, error) {
	if gRedis != nil {
		val, err := redis.String(gRedis.Do("get", key))
		if err != nil {
			return "", err
		}
		return val, nil
	}
	return "", errors.New("gRedis error, may not init...")
}
func DisConnRedis() {
	if gRedis != nil {
		gRedis.Close()
	}
}
