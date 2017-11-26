package args

import (
	"github.com/nourlikic/nond/database"
	"errors"
	"github.com/nourlikic/nond/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nourlikic/nond/accounts"
)

func ParseArgs(args []string) (err error) {

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
		return errors.New("no such usage")
	}
}

func argsAbi(args []string) (err error) {

	switch args[0] {
	case "load":
		if len(args) != 3 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		path := args[2]
		/*if b, _ := database.IsAbiExist(identifier); b {
			return util.ErrorAlreadyExistAbi
		} else {
			// TODO
		}*/
		if !util.IsExist(path) {
			return util.ErrorNoSuchFile
		}
		loadAbi(identifier, path)
		return nil
	case "list":
		if len(args) != 1 {
			return util.ErrorShowUsage
		}
		listAbi()
		return nil
	case "delete":
		if len(args) != 2 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		/*if b, e := database.IsAbiExist(identifier); !b {
			if e == nil {
				return util.ErrorNotExistAbi
			} else {
				return e
			}
		}*/
		deleteAbi(identifier)
		return nil
	case "detail":
		if len(args) != 2 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		/*if b, e := database.IsAbiExist(identifier); !b {
			if e == nil {
				return util.ErrorNotExistAbi
			} else {
				return e
			}
		}*/
		detailAbi(identifier)
		return nil
	default:
		return errors.New("show usage")
	}

}

func argsContract(args []string) (err error) {

	switch args[0] {
	case "txn":
		if len(args) < 3 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		method := args[2]
		transaction(identifier,method)
		return nil
	case "call":
		return nil
	case "save":
		if len(args) != 4 {
			return util.ErrorShowUsage
		}
		abiIdentifier := args[1]
		contractIdentifier := args[2]
		address := args[3]
		/*if b, e := database.IsAbiExist(abiIdentifier); !b {
			if e == nil {
				return util.ErrorNotExistAbi
			} else {
				return e
			}
		}
		if b, e := database.IsContractExist(contractIdentifier); b {
			fmt.Println(b)
			if e == nil {
				return util.ErrorAlreadyExistContract
			} else {
				return e
			}
		}*/
		//TODO is good address
		saveContract(contractIdentifier, address, abiIdentifier)
		return nil
	case "change":
		if len(args) != 3 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		address := args[2]
		/*if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return util.ErrorNotExistContract
			} else {
				return e
			}
		}*/
		//TODO is good address

		changeContractAddress(identifier, address)
		return nil
	case "delete":
		if len(args) != 2 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		/*if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return util.ErrorNotExistContract
			} else {
				return e
			}
		}*/
		deleteContract(identifier)
		return nil
	case "detail":
		if len(args) != 2 {
			return util.ErrorShowUsage
		}
		identifier := args[1]
		/*if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return util.ErrorNotExistContract
			} else {
				return e
			}
		}*/
		detailContract(identifier)
		return nil
	case "list":
		if len(args) != 1 {
			return util.ErrorShowUsage
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
		if len(args) != 1 {
			return util.ErrorShowUsage
		}
		newAccount()
		return nil
	case "delete":
		if len(args) != 3 {
			return util.ErrorShowUsage
		}
		address := common.HexToAddress(args[1])
		//TODO hide passphrase
		passphrase := args[2]
		if !accounts.IsExist(address) {
			return util.ErrorNotExistAccount
		}
		deleteAccount(address.Hex(), passphrase)
		return nil
	case "default":
		if len(args) != 2 {
			return util.ErrorShowUsage
		}
		address := common.HexToAddress(args[1])
		if !accounts.IsExist(address) {
			return util.ErrorNotExistAccount
		}
		defaultAccount(address.Hex())
		return nil
	case "list":
		if len(args) != 1 {
			return util.ErrorShowUsage
		}
		listAccounts()
		return nil
	default:
		return errors.New("show usage")
	}
}
