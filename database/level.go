package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/nourlikic/nond/config"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/nourlikic/nond/types"
	errors2 "github.com/nourlikic/nond/errors"
	"bytes"
	"fmt"
	"errors"
)

const (
	addressLenght = 40
)

var (
	contractPrefix = []byte("contract:")
	abiPrefix      = []byte("abi:")
	splitter       = []byte(":")
)

type LevelDatabase struct {
	fn       string
	database *leveldb.DB
}

func database(ldb *LevelDatabase) error {
	opts := &opt.Options{OpenFilesCacheCapacity: 5}
	if db, err := leveldb.OpenFile(config.GetDatabaseDir(), opts); err != nil {
		return err
	} else {
		ldb.database = db
		return nil
	}
}

func (db *LevelDatabase) get(key []byte) ([]byte, error) {
	if dat, err := db.database.Get(key, nil); err != nil {
		return nil, err
	} else {
		return dat, nil
	}
}

func (db *LevelDatabase) put(key []byte, value []byte) error {
	return db.database.Put(key, value, nil)
}

func (db *LevelDatabase) delete(key []byte) error {
	return db.database.Delete(key, nil)
}

func DetailContract(identifier string) (types.YContract, error) {

	var err error
	if address, abiIdentifier, err := getContract([]byte(identifier)); err == nil {
		if abi, err := getAbi(abiIdentifier); err == nil {
			contract := types.YContract{
				ContractIdentifier: identifier,
				AbiIdentifier:      string(abiIdentifier),
				Address:            string(address),
				Abi:                abi,
			}
			return contract, nil
		}
	}
	return types.YContract{}, err
}

func SaveContract(identifier string, address string, abi string) (err error) {
	return saveContract(identifier, address, abi)
}

func ListContracts() (contracts []types.YContract, err error) {
	return listContracts()
}

func DeleteContract(identifier string) (err error) {
	return deleteContract(identifier)
}

// ????
func IsContractExist(identifier string) (bool, error) {

	// TODO check err
	_, _, err := getContract([]byte(identifier))
	return isExist(err)
}

func SaveAbi(identifier string, abi []byte) (err error) {
	return saveAbi(identifier, abi)
}

func ListAbis() (abis []string, err error) {
	return listAbis()
}
func GetAbi(identifier string) (abi []byte, err error) {
	return getAbi([]byte(identifier))
}

func IsAbiExist(identifier string) (bool, error) {
	_, err := getAbi([]byte(identifier))
	return isExist(err)
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
			deleteContract(contract.ContractIdentifier)
		}
	}

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return err
	}
	defer db.database.Close()

	key := makeAbiKey([]byte(identifier))
	fmt.Println(string(key))
	if err != nil {
		return err
	}
	return db.delete(key)
}

func getContract(identifier []byte) (address []byte, abiIdentifier []byte, err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return nil, nil, err
	}
	defer db.database.Close()
	prefix := append(contractPrefix[:], []byte(identifier)[:]...)
	prefix = append(prefix[:], splitter[:]...)
	iter := db.database.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	for iter.Next() {
		if bytes.Compare(prefix, iter.Key()[:len(prefix)]) == 0 {
			_, address, err = splitContractKey(iter.Key())
			return address, iter.Value(), nil
		}
	}
	return nil, nil, leveldb.ErrNotFound
}

func listContracts() (contracts []types.YContract, err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return nil, err
	}
	defer db.database.Close()

	iter := db.database.NewIterator(util.BytesPrefix([]byte(contractPrefix)), nil)
	defer iter.Release()
	for iter.Next() {
		if contract, address, err := splitContractKey(iter.Key()); err == nil {

			yContract := types.YContract{
				string(iter.Value()),
				string(contract),
				string(address),
				nil,
			}
			contracts = append(contracts, yContract)
		}

	}
	if err := iter.Error(); err != nil {
		return nil, err
	}
	return contracts, nil
}

func saveContract(identifier string, address string, abi string) (err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return nil
	}
	defer db.database.Close()

	key := makeContractKey([]byte(identifier), []byte(address))
	return db.put(key, []byte(abi))
}

// TODO check
func getContractKey(identifier []byte) (key []byte, err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return nil, err
	}
	defer db.database.Close()
	prefix := append(contractPrefix[:], []byte(identifier)[:]...)
	prefix = append(prefix[:], splitter[:]...)
	iter := db.database.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	for iter.Next() {
		if bytes.Compare(prefix, iter.Key()[:len(prefix)]) == 0 {
			return iter.Key(), nil
		}
	}
	return nil, errors2.ErrorNotExistContract
}

func deleteContract(identifier string) (err error) {

	key, err := getContractKey([]byte(identifier))
	if err != nil {
		fmt.Println(err)
		return err
	}
	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return err
	}
	defer db.database.Close()

	if err != nil {
		return err
	}
	return db.delete(key)
}

func saveAbi(identifier string, abi []byte) (err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return err
	}
	defer db.database.Close()

	key := append(abiPrefix[:], []byte(identifier)[:]...)
	err = db.put(key, abi)
	db.database.Close()
	return err
}

func getAbi(identifier []byte) (abi []byte, err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return nil, err
	}
	defer db.database.Close()
	key := append(abiPrefix[:], []byte(identifier)[:]...)
	abi, err = db.get(key)
	if err != nil {
		return nil, err
	}
	return abi, nil
}

func listAbis() (abis []string, err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return nil, err
	}
	defer db.database.Close()
	iter := db.database.NewIterator(util.BytesPrefix([]byte(abiPrefix)), nil)
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

func splitContractKey(key []byte) (identifier []byte, address []byte, err error) {

	if !bytes.HasPrefix(key, contractPrefix) {
		return nil, nil, errors.New("No good prefix for contract")
	}

	keyLength := len(key)
	identifier = key[len(contractPrefix):keyLength-addressLenght-1]
	address = key[keyLength-addressLenght:]
	return identifier, address, nil
}

func makeContractKey(identifier []byte, address []byte) (key []byte) {
	key = append(contractPrefix, identifier[:]...)
	key = append(key[:], splitter[:]...)
	return append(key[:], address[:]...)
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

func isExist(err error) (bool, error) {

	if err == leveldb.ErrNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return 0 != 0, err
	}
}

/*
func ChangeContractAddress(identifier string, newAddress string) (err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return err
	}
	defer db.database.Close()

	prefix := append(contractPrefix[:], []byte(identifier)[:]...)
	iter := db.database.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	for iter.Next() {

		_identifier := []byte(identifier)
		prefix = append(prefix[:], _identifier[:]...)
		if bytes.Compare(prefix, iter.Key()[len(prefix):]) == 0 {
			if db.delete(iter.Key()) == nil {
				return db.put(makeContractKey(_identifier, []byte(newAddress)), iter.Value())
			}
		}
	}
	if err := iter.Error(); err != nil {
		return err
	}
	return err
}*/

func Iter() {

	db := LevelDatabase{}
	if err := database(&db); err != nil {
		fmt.Print(err)
		return
	}

	defer db.database.Close()
	iter := db.database.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Println("---")
		fmt.Println(string(iter.Key()))
		fmt.Println(string(iter.Value()))
	}
	iter.Release()
}

func IterFromTo(from string, to string) (err error) {

	db := LevelDatabase{}
	if err = database(&db); err != nil {
		return err
	}

	defer db.database.Close()
	iter := db.database.NewIterator(&util.Range{Start: []byte(from), Limit: []byte(to)}, nil)
	for iter.Next() {
		fmt.Println(string(iter.Key()) + " ---> " + string(iter.Value()))
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return err
	}
	return nil
}
