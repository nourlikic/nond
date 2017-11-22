package main

import (
	"github.com/nourlikic/nond/util"
	"github.com/nourlikic/nond/config"
)

func isInitialized() bool {

	return util.IsExist(config.GetAppDir()) &&
		util.IsExist(config.GetConfigPath()) &&
		util.IsExist(config.GetKeystoreDir()) &&
		util.IsExist(config.GetAbiDir()) &&
		util.IsExist(config.GetDatabaseDir())
}

func initialize() {

	if !util.IsExist(config.GetAppDir()) {
		util.MakeDir(config.GetAppDir(), 0700)
	}

	if !util.IsExist(config.GetConfigPath()) {
		config.CreateDefaultConfig()
	}

	if !util.IsExist(config.GetKeystoreDir()) {
		util.MakeDir(config.GetKeystoreDir(), 0700)
	}

	if !util.IsExist(config.GetAbiDir()) {
		util.MakeDir(config.GetAbiDir(), 0700)
	}

	if !util.IsExist(config.GetDatabaseDir()) {
		util.MakeDir(config.GetDatabaseDir(), 0700)
	}
}
