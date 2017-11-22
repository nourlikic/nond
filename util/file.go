package util

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"os/user"
)

func CreateFile(path string) {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if _, err := os.Create(path); err != nil {
			fmt.Println(err)
		}
	}
}

func MakeDir(dir string, perm os.FileMode) {

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, perm); err != nil {
			fmt.Println(err)
		}
	} else {
		log.Fatal("Directory Exists")
	}
}

func ReadJson(path string, v interface{}) error {

	var err error
	jsonFile, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonFile, v)
	if err != nil {
		return err
	}
	return nil
}

func WriteJson(path string, v interface{}) error {

	if bytes, err := json.MarshalIndent(v, "", "\t"); err != nil {
		return err
	} else {
		return ioutil.WriteFile(path, bytes, 0700)
	}

}

func ReadFile(path string) ([]byte,error){

	return ioutil.ReadFile(path)
}

func GetHomeDir() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func IsExist(dir string) bool {

	if _, err := os.Stat(dir); err == nil {
		return true
	}
	return false
}
