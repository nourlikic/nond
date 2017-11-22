package accounts

import (
	"fmt"
	"github.com/nourlikic/nond/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nourlikic/nond/errors"
	"github.com/ethereum/go-ethereum/accounts"
)

func CreateNewAccount(passphrase string) error {

	account, err := createNewAccount(passphrase)
	if err != nil {
		return err
	}
	fmt.Println(account.Address.Hex())
	return nil
}

func ListAccounts() error {

	accounts := listAccounts()
	c, err := config.GetConfig()
	if err != nil {
		return err
	}
	defaultAddress := common.HexToAddress(c.DefaultAccount)
	for _, account := range accounts {
		fmt.Print(account.Address.Hex())
		if account.Address.Hex() == defaultAddress.Hex() {
			fmt.Print("  Default")
		}
		fmt.Println("")
	}
	return nil
}

func DefaultAccount(addr string) error {

	defaultAddress := common.HexToAddress(addr)
	for _, account := range listAccounts() {
		if account.Address.Hex() == defaultAddress.Hex() {
			conf, err := config.GetConfig()
			if err != nil {
				return err
			}
			conf.DefaultAccount = defaultAddress.Hex()
			config.UpdateConfig(conf)
			return nil
		}
	}
	return errors.ErrorNotExistAccount
}

/*func IsDefaultAccount(account accounts.Account) (bool,error) {
	acc,err := getDefaultAccount()
	if err!= nil{
		return false,err
	}
	if account == acc {
		return true,nil
	}
	return false,nil
}*/

func IsDefaultAccount(addr string) (bool, error) {
	address := common.HexToAddress(addr)
	acc, err := getDefaultAccount()
	if err != nil {
		return false, err
	}
	if address == acc.Address {
		return true, nil
	}
	return false, nil
}

func getDefaultAccount() (accounts.Account, error) {

	conf, err := config.GetConfig()
	if err != nil {
		return accounts.Account{}, err
	}
	defaultAddress := common.HexToAddress(conf.DefaultAccount)
	for _, account := range listAccounts() {
		if account.Address.Hex() == defaultAddress.Hex() {
			return account, nil
		}
	}
	return accounts.Account{}, errors.ErrorNotExistAccount
}

func DeleteAccount(addr string, passphrase string) error {

	address := common.HexToAddress(addr)
	for _, account := range listAccounts() {
		if account.Address.Hex() == address.Hex() {
			err := delete(account, passphrase)
			if err != nil {
				return err
			}
			conf, err := config.GetConfig()
			if err != nil {
				return err
			}
			b, err := IsDefaultAccount(conf.DefaultAccount)
			if err != nil {
				return err
			}
			if b {
				conf.DefaultAccount = ""
				config.UpdateConfig(conf)
			}
			return nil
		}
	}
	return errors.ErrorNotExistAccount
}
