package config

import "github.com/nourlikic/nond/util"

const (
	abiDir      = "/abi"
	databaseDir = "/database"
	keystoreDir  = "/keystore"
	configPath  = "/conf.json"
)

func GetAppDir() string {

	return util.GetHomeDir() + "/.nond"
}

func GetAbiDir() string {
	return GetAppDir() + abiDir
}

func GetDatabaseDir() string {
	return GetAppDir() + databaseDir
}

func GetConfigPath() string {

	return GetAppDir() + configPath
}

func GetKeystoreDir() string {

	return GetAppDir() + keystoreDir
}
