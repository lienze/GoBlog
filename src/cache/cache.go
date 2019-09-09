package cache

import "errors"

func InitCache(cacheType string) error {
	switch cacheType {
	case "redis":
		InitRedis()
	default:
		return errors.New("cache type error")
	}
	return nil
}
