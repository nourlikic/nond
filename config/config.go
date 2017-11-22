package config

import (
	"github.com/nourlikic/nond/util"
	"fmt"
)

type Config struct {
	RelayUrl       string `json:"relay_url"`
	DefaultAccount string `json:"default_account"`
	GasPrice       uint   `json:"gas_price"`
	Gas            uint   `json:"gas"`
}

func CreateDefaultConfig() {

	config := Config{
		RelayUrl : "localhost:8545",
		DefaultAccount: "",
		GasPrice:20000,
		Gas:90000,
	}
	err := UpdateConfig(config)
	fmt.Println(err)
}

func UpdateConfig(config Config) error{

	return util.WriteJson(GetConfigPath(),config)
}

func GetConfig() (Config,error) {

	var config Config
	err := util.ReadJson(GetConfigPath(), &config)
	return config,err
}
