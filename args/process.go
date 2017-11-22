package args
/*
import (
	"fmt"
	"errors"
	"github.com/nourlikic/nond/util"
	"github.com/nourlikic/nond/database"
	"github.com/nourlikic/nond/accounts"
	"github.com/ethereum/go-ethereum/common"
	errors2 "github.com/nourlikic/nond/errors"
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

func ProcessArgs(args []string) (action Action, _args map[string]string, err error) {

	switch args[0] {
	case "abi":
		return argsAbi(args[1:])
	case "contract":
		return argsContract(args[1:])
	case "account":
		return argsAccount(args[1:])
	case "iter":
		return ActionIter,nil,nil
	default:
		return "", nil, errors.New("show usage")
	}
}

func Usage() {

}

// load,list,delete
func argsAbi(args []string) (action Action, _args map[string]string, err error) {

	switch args[0] {
	case "load":
		if len(args) != 3 {
			return "", nil, errors2.ErrorShowUsage
		}
		identifier := args[1]
		path := args[2]
		if b, _ := database.IsAbiExist(identifier); b {
			return "", nil, errors2.ErrorAlreadyExistAbi
		} else {
			// TODO
		}
		if !util.IsExist(path) {
			return "", nil, errors2.ErrorNoSuchFile
		}
		_args = map[string]string{
			ArgAbiIdentifier: identifier,
			ArgPath:          path,
		}
		return ActionLoadAbi, _args, nil
	case "list":
		if len(args) != 1 {
			return "", nil, errors2.ErrorShowUsage
		}
		return ActionListAbi, nil, nil
	case "delete":
		if len(args) != 2 {
			return "", nil, errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsAbiExist(identifier); !b {
			if e == nil {
				return "", nil, errors2.ErrorNotExistAbi
			} else {
				return "", nil, e
			}
		}
		_args = map[string]string{
			ArgAbiIdentifier: identifier,
		}
		return ActionDeleteAbi, _args, nil
	case "detail":
		if len(args) != 2{
			return "", nil, errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsAbiExist(identifier); !b {
			if e == nil {
				return "", nil, errors2.ErrorNotExistAbi
			} else {
				return "", nil, e
			}
		}
		_args = map[string]string{
			ArgAbiIdentifier: identifier,
		}
		return ActionDetailAbi, _args, nil
	default:
		return "", nil, errors.New("show usage")
	}

}

func argsContract(args []string) (action Action, _args map[string]string, err error) {

	switch args[0] {
	case "save":
		if len(args) != 4 {
			return "", nil, errors2.ErrorShowUsage
		}
		abiIdentifier := args[1]
		contractIdentifier := args[2]
		address := args[3]
		if b, e := database.IsAbiExist(abiIdentifier); !b {
			if e == nil {
				return "", nil, errors2.ErrorNotExistAbi
			} else {
				return "", nil, e
			}
		}
		if b, e := database.IsContractExist(abiIdentifier); b {
			fmt.Println(b)
			if e == nil {
				return "", nil, errors2.ErrorAlreadyExistContract
			} else {
				return "", nil, e
			}
		}
		//TODO is good address
		_args = map[string]string{
			ArgAbiIdentifier:      abiIdentifier,
			ArgContractIdentifier: contractIdentifier,
			ArgAddress:            address,
		}
		return ActionSaveContract, _args, nil
	case "change":
		if len(args) != 3 {
			return "", nil, errors2.ErrorShowUsage
		}
		identifier := args[1]
		address := args[2]
		if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return "", nil, errors2.ErrorNotExistContract
			} else {
				return "", nil, e
			}
		}
		//TODO is good address
		_args = map[string]string{
			ArgContractIdentifier: identifier,
			ArgAddress:            address,
		}
		return ActionChangeContractAddress, _args, nil
	case "delete":
		if len(args) != 2 {
			return "", nil, errors2.ErrorShowUsage
		}
		identifier := args[1]
		if b, e := database.IsContractExist(identifier); !b {
			if e == nil {
				return "", nil, errors2.ErrorNotExistContract
			} else {
				return "", nil, e
			}
		}
		_args = map[string]string{
			ArgContractIdentifier: identifier,
		}
		return ActionDeleteContract, _args, nil
	case "list":
		if len(args) != 1 {
			return "", nil, errors2.ErrorShowUsage
		}
		return ActionListContracts, nil, nil
	default:
		return "", nil, errors.New("show usage")
	}
}

func argsAccount(args []string) (action Action, _args map[string]string, err error) {

	switch args[0] {
	case "new":
		if len(args) != 2 {
			return "", nil, errors2.ErrorShowUsage
		}
		//TODO hide passphrase
		passphrase := args[1]
		_args = map[string]string{
			ArgPassphrase: passphrase,
		}
		return ActionNewAccount, _args, nil
	case "delete":
		if len(args) != 3 {
			return "", nil, errors2.ErrorShowUsage
		}
		address := common.HexToAddress(args[1])
		//TODO hide passphrase
		passphrase := args[2]
		if !accounts.IsExist(address) {
			return "", nil, errors2.ErrorNotExistAccount
		}
		_args = map[string]string{
			ArgAddress: address.Hex(),
			ArgAddress: address.Hex(),
		}
		return ActionDeleteAccount, _args, nil
	case "default":
		if len(args) != 2 {
			return "", nil, errors2.ErrorShowUsage
		}
		address := common.HexToAddress(args[1])
		if !accounts.IsExist(address) {
			return "", nil, errors2.ErrorNotExistAccount
		}
		_args = map[string]string{
			ArgAddress: address.Hex(),
		}
		return ActionDefaultAccount, _args, nil
	case "list":
		if len(args) != 1 {
			return "", nil, errors2.ErrorShowUsage
		}
		return ActionListAccounts, nil, nil
	default:
		return "", nil, errors.New("show usage")
	}
}

func usage() {

	fmt.Println("No args!")
}
*/