package db

import (
	"GoBlog/src/config"

	"github.com/lienze/go2db/dao"
)

func InitMongo() {
	dao.InitDB(config.GConfig.DB.DBName)
}
