package db

import (
	"GoBlog/src/config"
	"fmt"
)

func InitDB() error {
	fmt.Println("InitDB...", config.GConfig.DB.DBType)
	dbtype := config.GConfig.DB.DBType
	switch dbtype {
	case "mongodb":
		return InitMongo()
	case "redis":
		return InitRedis()
	default:
		return fmt.Errorf("Can not init DB Type %s", dbtype)
	}
}
