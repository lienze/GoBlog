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

func DisConnRedis() {
	if gRedis != nil {
		gRedis.Close()
	}
}

//------------------------------------------------------------------------string
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

//--------------------------------------------------------------------------set
func AddSetKeyValue(key string, val string) error {
	if gRedis != nil {
		_, err := redis.String(gRedis.Do("sadd", key, val))
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("gRedis error, may not init...")
}

func GetSetSize(key string) (int, error) {
	if gRedis != nil {
		setsize, err := redis.Int(gRedis.Do("scard", key))
		if err != nil {
			return -1, err
		}
		return setsize, nil
	}
	return -1, errors.New("gRedis error, may not init...")
}

func PrintSet(key string) {
	if gRedis != nil {
		ret, err := redis.Strings(gRedis.Do("smembers", key))
		if err != nil {
			fmt.Println("PrintSet Error:", err)
			return
		}
		fmt.Println("PrintSet:", key, ret)
	}
}
