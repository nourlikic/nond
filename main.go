package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/nourlikic/nond/util"
	"github.com/nourlikic/nond/contracts"
	"github.com/nourlikic/nond/database"
	"github.com/nourlikic/nond/accounts"
	"github.com/ethereum/go-ethereum/common"
	errors2 "github.com/nourlikic/nond/errors"
)

const (
	ActionRefresh               Action = "Refresh"
	ActionLoadAbi               Action = "LoadAbi"              // identity, path
	ActionListAbi               Action = "ListAbi"              // nil
	ActionDeleteAbi             Action = "DeleteAbi"            // identifier
	ActionDetailAbi             Action = "DetailAbi"            // identifier
	ActionSaveContract          Action = "SaveContract"         // address, abi
	ActionDeleteContract        Action = "DeleteContract"       // identifier
	ActionListContracts         Action = "ListContracts"        // nil
	ActionDetailContract        Action = "DetailContract"       // identifier
	ActionChangeContractAddress Action = "ChangeContractAddres" // identifier, address
	ActionNewAccount            Action = "NewAccount"           // identifier,passphrase
	ActionDeleteAccount         Action = "DeleteAccount"        // identifier
	ActionDefaultAccount        Action = "DefaultAccount"       //identifier
	ActionListAccounts          Action = "ListAccount"          // nil

	ActionCall           Action = "Call"        // cidentifier, args
	ActionTransaction    Action = "Transaction" // cidentifier, args
	ActionTestConnection Action = "TestConnection"
	ActionRelayUrl       Action = "SetRelayUrl"

	ActionChageGasPrice Action = "ChangeGasPrice"
	ActionChangeGas     Action = "ChangeGas"

	//Temp
	ActionIter Action = "iter"
)

type Action string
type Arg string

const (
	ArgContractIdentifier string = "cid"
	ArgAbiIdentifier      string = "aid"
	ArgPath               string = "path"
	ArgAddress            string = "address"
	ArgPassphrase         string = "passphrase"
)

func processArgs(args []string) (err error) {

	switch args[0] {
	case "abi":
		return argsAbi(args[1:])
	case "contract":
		return argsContract(args[1:])
	case "account":
		return argsAccount(args[1:])
	case "iter":
		database.Iter()
		return nil
	default:
		return errors.New("show usage")
	}
}

func argsAbi(args []string) (err error) {

	switch args[0] {
	case "load":
		if len(args) != 3 {
			return errors2.ErrorShowUsage
		}
		identifier := args[1]
		path := args[2]
		if b, _ := database.IsAbiExist(identifier); b {
			return errors2.ErrorAlreadyExistAbi
		} else {
			// TODO
		}
		if !util.IsExist(path) {
			return errors2.ErrorNoSuchFile
		}
		loadAbi(identifier, path)
		return nil
	case "list":
		if len(args) != 1 {
			return errors2.ErrorShowUsage
		}
		listAbi()
		return nil
	case "delete":
		if len(args) != 2 {
			return errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsAbiExist(identifier); !b {
			if e == nil {
				return errors2.ErrorNotExistAbi
			} else {
				return e
			}
		}
		deleteAbi(identifier)
		return nil
	case "detail":
		if len(args) != 2 {
			return errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsAbiExist(identifier); !b {
			if e == nil {
				return errors2.ErrorNotExistAbi
			} else {
				return e
			}
		}
		detailAbi(identifier)
		return nil
	default:
		return errors.New("show usage")
	}

}

func argsContract(args []string) (err error) {

	switch args[0] {
	case "save":
		if len(args) != 4 {
			return errors2.ErrorShowUsage
		}
		abiIdentifier := args[1]
		contractIdentifier := args[2]
		address := args[3]
		if b, e := database.IsAbiExist(abiIdentifier); !b {
			if e == nil {
				return errors2.ErrorNotExistAbi
			} else {
				return e
			}
		}
		if b, e := database.IsContractExist(abiIdentifier); b {
			fmt.Println(b)
			if e == nil {
				return errors2.ErrorAlreadyExistContract
			} else {
				return e
			}
		}
		//TODO is good address
		saveContract(contractIdentifier,address,abiIdentifier)
		return nil
	case "change":
		if len(args) != 3 {
			return errors2.ErrorShowUsage
		}
		identifier := args[1]
		address := args[2]
		if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return errors2.ErrorNotExistContract
			} else {
				return e
			}
		}
		//TODO is good address

		changeContractAddress(identifier,address)
		return nil
	case "delete":
		if len(args) != 2 {
			return errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return errors2.ErrorNotExistContract
			} else {
				return e
			}
		}
		deleteContract(identifier)
		return nil
	case "detail":
		if len(args) != 2 {
			return errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return errors2.ErrorNotExistContract
			} else {
				return e
			}
		}
		detailContract(identifier)
		return nil
	case "list":
		if len(args) != 1 {
			return errors2.ErrorShowUsage
		}
		listContracts()
		return nil
	default:
		return errors.New("show usage")
	}
}

func argsAccount(args []string) (err error) {

	switch args[0] {
	case "new":
		if len(args) != 2 {
			return errors2.ErrorShowUsage
		}
		//TODO hide passphrase
		passphrase := args[1]
		newAccount(passphrase)
		return nil
	case "delete":
		if len(args) != 3 {
			return errors2.ErrorShowUsage
		}
		address := common.HexToAddress(args[1])
		//TODO hide passphrase
		passphrase := args[2]
		if !accounts.IsExist(address) {
			return errors2.ErrorNotExistAccount
		}
		deleteAccount(address.Hex(),passphrase)
		return nil
	case "default":
		if len(args) != 2 {
			return errors2.ErrorShowUsage
		}
		address := common.HexToAddress(args[1])
		if !accounts.IsExist(address) {
			return errors2.ErrorNotExistAccount
		}
		defaultAccount(address.Hex())
		return nil
	case "list":
		if len(args) != 1 {
			return errors2.ErrorShowUsage
		}
		listAccounts()
		return nil
	default:
		return errors.New("show usage")
	}
}

func main() {

	if !isInitialized() {
		fmt.Println("Will Initialize")
		initialize()
	}

	if len(os.Args) < 2 {
		fmt.Println("hi :)")
		os.Exit(0)
	}

	err := processArgs(os.Args[1:])
	if err != nil {
		fmt.Println("ERR")
		fmt.Println(err)
		os.Exit(0)
	}

	//addr := common.HexToAddress("1d8bb0334e741d8b82ba05260cc6b0f9d7ace04e")
	//database.SaveContract("some_contract",addr.String(),"some_abi")
}

func loadAbi(identifier string, path string) {
	err := contracts.LoadAbi(identifier, path)
	if err != nil {
		fmt.Println(err)
	}
}

func listAbi() {
	err := contracts.ListAbi()
	if err != nil {
		fmt.Println(err)
	}
}

func detailAbi(identifier string) {
	err := contracts.DetailAbi(identifier)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteAbi(identifier string) {
	err := contracts.DeleteAbi(identifier)
	if err != nil {
		fmt.Println(err)
	}
}

func saveContract(contractIdentifier string, address string, abiIdenfier string) {
	err := contracts.SaveContract(contractIdentifier, address, abiIdenfier)
	if err != nil {
		fmt.Println(err)
	}
}

func detailContract(identifier string) {
	err := contracts.DetailContract(identifier)
	if err != nil {
		fmt.Println(err)
	}
}

func listContracts() {
	err := contracts.ListContracts()
	if err != nil {
		fmt.Println(err)
	}
}

func deleteContract(identifier string) {
	err := contracts.DeleteContract(identifier)
	if err != nil {
		fmt.Println(err)
	}
}

func changeContractAddress(identifier string,address string) {

}

func newAccount(passphrase string) {
	err := accounts.CreateNewAccount(passphrase)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteAccount(address string, passphrase string) {
	err := accounts.DeleteAccount(address, passphrase)
	if err != nil {
		fmt.Println(err)
	}
}

func defaultAccount(address string) {
	err := accounts.DefaultAccount(address)
	if err != nil {
		fmt.Println(err)
	}
}

func listAccounts() {
	err := accounts.ListAccounts()
	if err != nil {
		fmt.Println(err)
	}
}
