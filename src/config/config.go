package config

import (
	"github.com/BurntSushi/toml"
)

var GConfig tomlConfig

func InitConfig() {
	filePath := "./config/global.toml"
	if _, err := toml.DecodeFile(filePath, &GConfig); err != nil {
		panic(err)
	}
	//fmt.Println(GConfig)
}
