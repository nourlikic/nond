package contracts

import (
	"fmt"
	"math/big"
	"strconv"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/nourlikic/nond/accounts"
	"github.com/nourlikic/nond/net"
	"github.com/nourlikic/nond/config"
	"github.com/ethereum/go-ethereum/core/types"
)

func Transaction(contractIdentifier string, methodName string,jsonAccount []byte) ([]byte,error){

	address, abiIdentifier, err := GetContract(contractIdentifier)

	if err != nil {
		return nil,err
	}
	abi, err := GetAbiObject(abiIdentifier)
	if err != nil {
		return nil,err
	}
	var method ethabi.Method
	b := false
	for _, m := range abi.Methods {
		if methodName == m.Name {
			method = m
			b = true
			break
		}
	}
	if !b {return nil,errors.New("No such method")}

	data,err := makeInputData(abi,method)
	if err != nil {
		return nil,err
	}

	tOpts,err := makeTransactor()
	if err != nil {
		return nil,err
	}
	signedTxn,err := signTransaction(tOpts,address,data)
	if err != nil {
		fmt.Println("BAD TXN")
		fmt.Println(err)
	}
	var txns types.Transactions
	txns = append(txns, signedTxn)
	//return fmt.Sprintf("%x",txns.GetRlp(0))),nil
	return txns.GetRlp(0),nil

}

func signTransaction(
	opts *bind.TransactOpts,
	address common.Address,
	data []byte) (*types.Transaction, error) {

	rawTx := types.NewTransaction(
		opts.Nonce.Uint64(),
		address,
		opts.Value,
		opts.GasLimit,
		opts.GasPrice,
		data)
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func makeInputData(abi ethabi.ABI,m ethabi.Method) ([]byte,error){

	inputs := scanUserInputs(m)
	parsedInputs, err := parseInputs(inputs)
	if err != nil {
		return nil, err
	}
	data,err := abi.Pack(m.Name,parsedInputs...)
	if err != nil {
		return nil,err
	}
	return data,nil
}

func makeTransactor() (*bind.TransactOpts,error){

	passphrase := accounts.ScanPassphrase()
	defAcc, err := accounts.ReadEncDefaultAccount()
	if err != nil {
		return nil,err
	}

	tOpts, err := bind.NewTransactor(strings.NewReader(string(defAcc)), passphrase)
	if err != nil {
		return nil,err
	}
	tOpts.Nonce = net.GetNonce()
	tOpts.Value = big.NewInt(0)
	config, err := config.GetConfig()
	if err != nil {
		return nil,err
	}
	tOpts.GasPrice = big.NewInt(int64(config.GasPrice))
	tOpts.GasLimit = big.NewInt(int64(config.Gas))
	return tOpts,nil
}

func scanUserInputs(method ethabi.Method) map[ethabi.Argument]string {

	inputs := make(map[ethabi.Argument]string)
	fmt.Println(method.String())
	for _, argument := range method.Inputs {
		var str string
		fmt.Println("Enter value for, " + argument.Name + ", " + argument.Type.String())
		fmt.Scanf("%s", &str)
		inputs[argument] = str
	}

	return inputs
}

func parseInputs(inputs map[ethabi.Argument]string) (params []interface{}, err error) {

	for argument, input := range inputs {

		typ := argument.Type.String()
		size := argument.Type.Size
		if typ[:3] == "int" {

			if size > 64 {
				i, err := strconv.ParseInt(input, 10, 64)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Cannot parse %s, %s", argument.Name, input))
				}
				params = append(params, big.NewInt(i))
			} else {
				i, err := strconv.ParseInt(input, 10, size)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Cannot parse %s, %s", argument.Name, input))
				}
				if size == 8 {
					params = append(params, int8(i))
				} else if size == 16 {
					params = append(params, int16(i))
				} else if size == 32 {
					params = append(params, int32(i))
				} else if size == 64 {
					params = append(params, int64(i))
				} else {
					return nil, errors.New(fmt.Sprintf("Cannot parse, size: %d, type: %s", size, argument.Type.String()))
				}
			}
		} else if typ[:4] == "uint" {

			if size > 64 {
				i, err := strconv.ParseInt(input, 10, 64)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Cannot parse %s, %s", argument.Name, input))
				}
				params = append(params, big.NewInt(i))
			} else {

				i, err := strconv.ParseUint(input, 10, size)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Cannot parse %s, %s", argument.Name, input))
				}
				if size == 8 {
					params = append(params, uint8(i))
				} else if size == 16 {
					params = append(params, uint16(i))

				} else if size == 32 {
					params = append(params, uint32(i))

				} else if size == 64 {
					params = append(params, uint64(i))
				} else {
					return nil, errors.New(fmt.Sprintf("Cannot parse, size: %d, type: %s", size, argument.Type.String()))
				}
			}
		} else if typ == "bool" {
			i, err := strconv.ParseBool(input)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Cannot parse %s, %s", argument.Name, input))
			}
			params = append(params, i)
		} else if typ == "address" {
			if common.IsHexAddress(input) {
				address := common.HexToAddress(input)
				params = append(params, address)
			} else {
				return nil, errors.New(fmt.Sprintf("Cannot parse %s, %s", argument.Name, input))
			}
		} else if typ == "string" {
			params = append(params, input)
		} else if typ[:5] == "bytes" {
			if size > 0 {
				var arr [32]byte
				//arr := make([]byte, size)
				copy(arr[:], input)
				params = append(params, arr)
			} else {
				//arr := make([]byte, len(input))
				var arr [32]byte
				copy(arr[:], input)
				params = append(params, arr)
			}

		}
	}

	return params, nil

}
