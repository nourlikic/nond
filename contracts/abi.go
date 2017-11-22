package contracts

import (
	"github.com/nourlikic/nond/util"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/nourlikic/nond/database"
)

//import "github.com/ethereum/go-ethereum/contracts/ens/contract"

func LoadAbi(identifier string, path string) (err error) {

	jsondata, err := util.ReadFile(path)
	if err != nil {
		return err
	}

	_abi, err := abi.JSON(strings.NewReader(string(jsondata)))
	if err != nil {
		return err
	}

	err = database.SaveAbi(identifier, jsondata)
	if err != nil {
		return err
	}

	for _, m := range _abi.Methods {
		fmt.Println(m.String())
	}

	return nil
}

func DetailAbi(identifier string) (err error) {

	abidata, err := database.GetAbi(identifier)
	abi, err := abi.JSON(strings.NewReader(string(abidata)))
	if err != nil {
		return err
	}
	for _, m := range abi.Methods {
		fmt.Println(m.String())
	}

	return nil
}

func ListAbi() (err error) {

	abis, err := database.ListAbis()
	if err != nil {
		return err
	}
	for _, abi := range abis {
		fmt.Println(abi)
	}
	return nil
}

func DeleteAbi(identifier string) (err error) {

	err = database.DeleteAbi(identifier)
	if err != nil {
		return err
	}
	return nil
}
