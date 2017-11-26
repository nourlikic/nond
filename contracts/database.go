package contracts

import (
	"fmt"
	"bytes"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nourlikic/nond/util"
	"github.com/nourlikic/nond/database"
	"errors"
	"strings"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/nourlikic/nond/types"
)

var (
	contractPrefix = []byte("contract:")
	abiPrefix      = []byte("abi:")
	splitter       = []byte(":")
)

func ListContracts() (contracts []types.Contract, err error) {
	return listContracts()
}

func SaveContract(contractIdentifier string, address string, abiIdentifier string) error {
	if common.IsHexAddress(address) {
		return saveContract([]byte(contractIdentifier), common.HexToAddress(address).Bytes(), []byte(abiIdentifier))
	}
	return util.ErrorBadAddress
}

func DeleteContract(identifier string) error {
	return deleteContract([]byte(identifier))
}

func GetContract(identifier string) (common.Address, string, error) {
	address, abiIdentifier, err := getContract([]byte(identifier))
	if err == leveldb.ErrNotFound {
		return common.Address{}, "", util.ErrorNotExistContract
	}
	return common.BytesToAddress(address), string(abiIdentifier), nil
}

func getContract(identifier []byte) (address []byte, abiIdentifier []byte, err error) {
	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()
	prefix := append(contractPrefix[:], []byte(identifier)[:]...)
	prefix = append(prefix[:], splitter[:]...)
	iter := db.Iterator(prefix)
	defer iter.Release()
	for iter.Next() {
		if bytes.Compare(prefix, iter.Key()[:len(prefix)]) == 0 {
			_, address, err = splitContractKey(iter.Key())
			return address, iter.Value(), nil
		}
	}
	return nil, nil, leveldb.ErrNotFound
}

func listContracts() (_contracts []types.Contract, err error) {
	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	iter := db.Iterator([]byte(contractPrefix))
	defer iter.Release()
	for iter.Next() {
		if identifier, addr, err := splitContractKey(iter.Key()); err == nil {

			address := common.BytesToAddress(addr)

			c := types.Contract{
				AbiIdentifier:      string(iter.Value()),
				ContractIdentifier: string(identifier),
				Address:            address,
			}
			_contracts = append(_contracts, c)
		}

	}
	if err := iter.Error(); err != nil {
		return nil, err
	}
	return _contracts, nil
}

func saveContract(identifier []byte, address []byte, abi []byte) (err error) {
	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return err
	}
	defer db.Close()
	key := makeContractKey(identifier, address)
	return db.Put(key, abi)
}

// TODO check
func getContractKey(identifier []byte) (key []byte, err error) {
	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	prefix := append(contractPrefix[:], []byte(identifier)[:]...)
	prefix = append(prefix[:], splitter[:]...)
	iter := db.Iterator(prefix)
	defer iter.Release()
	for iter.Next() {
		if bytes.Compare(prefix, iter.Key()[:len(prefix)]) == 0 {
			return iter.Key(), nil
		}
	}
	return nil, util.ErrorNotExistContract
}

func deleteContract(identifier []byte) (err error) {
	key, err := getContractKey(identifier)
	if err != nil {
		fmt.Println(err)
		return err
	}
	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	if err != nil {
		return err
	}
	return db.Delete(key)
}

func splitContractKey(key []byte) (identifier []byte, address []byte, err error) {
	if !bytes.HasPrefix(key, contractPrefix) {
		return nil, nil, errors.New("No good prefix for contract")
	}
	keyLength := len(key)
	identifier = key[len(contractPrefix):keyLength-common.AddressLength-1]
	address = key[keyLength-common.AddressLength:]
	return identifier, address, nil
}

func makeContractKey(identifier []byte, address []byte) (key []byte) {
	key = append(contractPrefix, identifier[:]...)
	key = append(key[:], splitter[:]...)
	return append(key[:], address[:]...)
}


/*

Abi Database
*/


func LoadAbi(identifier string, path string) (ethabi.ABI, error) {

	jsondata, err := util.ReadFile(path)
	if err != nil {
		return ethabi.ABI{}, err
	}

	_abi, err := ethabi.JSON(strings.NewReader(string(jsondata)))
	if err != nil {
		return ethabi.ABI{}, err
	}

	err = saveAbi(identifier, jsondata)
	if err != nil {
		return ethabi.ABI{}, err
	}

	return _abi, nil
}

func GetAbiObject(identifier string) (ethabi.ABI, error) {

	abidata, err := getAbi([]byte(identifier))
	abi, err := ethabi.JSON(strings.NewReader(string(abidata)))
	if err != nil {
		return ethabi.ABI{}, err
	}
	return abi, nil
}

func ListAbis() ([]string, error) {
	return listAbis()
}

func DeleteAbi(identifier string) (err error) {
	return deleteAbi(identifier)
}

func deleteAbi(identifier string) (err error) {

	contracts, err := listContracts()
	if err != nil {
		return err
	}
	for _, contract := range contracts {

		if contract.AbiIdentifier == identifier {
			deleteContract([]byte(contract.ContractIdentifier))
		}
	}

	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	key := makeAbiKey([]byte(identifier))
	fmt.Println(string(key))
	if err != nil {
		return err
	}
	return db.Delete(key)
}

func saveAbi(identifier string, abi []byte) (err error) {

	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	key := append(abiPrefix[:], []byte(identifier)[:]...)
	err = db.Put(key, abi)
	return err
}

func getAbi(identifier []byte) (abi []byte, err error) {

	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return nil,err
	}
	defer db.Close()
	key := append(abiPrefix[:], []byte(identifier)[:]...)
	abi, err = db.Get(key)
	if err != nil {
		return nil, err
	}
	return abi, nil
}

func listAbis() (abis []string, err error) {

	db := database.LevelDB{}
	err = db.Init()
	if err != nil {
		return nil,err
	}
	defer db.Close()
	iter := db.Iterator(abiPrefix)
	defer iter.Release()
	for iter.Next() {
		if identifier, err := splitAbiKey(iter.Key()); err == nil {
			abis = append(abis, string(identifier))
		}
	}
	if err := iter.Error(); err != nil {
		panic(err)
	}
	return abis, nil
}

func makeAbiKey(identifier []byte) (key []byte) {

	return append(abiPrefix, identifier[:]...)
}

func splitAbiKey(key []byte) (identifier []byte, err error) {

	if !bytes.HasPrefix(key, abiPrefix) {
		return nil, errors.New("No good prefix for abi")
	}
	return key[len(abiPrefix):], nil
}