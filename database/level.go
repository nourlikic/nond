package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/nourlikic/nond/config"
	"github.com/syndtr/goleveldb/leveldb/util"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type LevelDB struct {
	fn       string
	database *leveldb.DB
}

func database(ldb *LevelDB) error {
	opts := &opt.Options{OpenFilesCacheCapacity: 5}
	if db, err := leveldb.OpenFile(config.GetDatabaseDir(), opts); err != nil {
		return err
	} else {
		ldb.database = db
		return nil
	}
}

func (db *LevelDB) Iterator(prefix []byte) iterator.Iterator {

	return db.database.NewIterator(util.BytesPrefix(prefix), nil)
}

func (db *LevelDB) Init() (err error) {

	opts := &opt.Options{OpenFilesCacheCapacity: 5}
	db.database, err = leveldb.OpenFile(config.GetDatabaseDir(), opts)
	return err
}

func (db *LevelDB) Close() {
	db.database.Close()
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	if dat, err := db.database.Get(key, nil); err != nil {
		return nil, err
	} else {
		return dat, nil
	}
}

func (db *LevelDB) Put(key []byte, value []byte) error {
	return db.database.Put(key, value, nil)
}

func (db *LevelDB) Delete(key []byte) error {
	return db.database.Delete(key, nil)
}

func IsExist(err error) (bool, error) {

	if err == leveldb.ErrNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return 0 != 0, err
	}
}

// ????
/*func IsContractExist(identifier string) (bool, error) {

	// TODO check err
	_, _, err := getContract([]byte(identifier))
	return isExist(err)
}

func IsAbiExist(identifier string) (bool, error) {
	_, err := getAbi([]byte(identifier))
	return isExist(err)
}*/

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

func Iter() error{

	db := LevelDB{}
	err := db.Init()
	if err != nil {
		return err
	}
	defer db.Close()
	iter := db.database.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Println("---")
		fmt.Println(string(iter.Key()))
		fmt.Println(string(iter.Value()))
	}
	iter.Release()
	return nil
}