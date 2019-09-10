package cache

import "fmt"

func InitCache(cacheType string) error {
	switch cacheType {
	case "redis":
		return InitRedis()
	default:
		//return errors.New("cache type error")
		return fmt.Errorf("cache type error:%s", cacheType)
	}
}
