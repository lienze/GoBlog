package db

import (
	"GoBlog/src/config"
	"errors"

	"github.com/lienze/go2db/dao"
)

func InitMongo() error {
	bInit := dao.InitDB(config.GConfig.DB.DBName)
	if !bInit {
		return errors.New("init mongodb error")
	}
	return nil
}
