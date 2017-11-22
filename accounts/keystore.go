package accounts

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nourlikic/nond/config"
)

const (
	scryptN = keystore.StandardScryptN
	scryptP = keystore.StandardScryptP
)

func createNewAccount(passphrase string) (accounts.Account, error) {

	return getKeystore().NewAccount(passphrase)
}

func IsExist(address common.Address) bool {
	return getKeystore().HasAddress(address)
}

func listAccounts() []accounts.Account {
	return getKeystore().Accounts()
}

func delete(account accounts.Account, passphrase string) error {
	return getKeystore().Delete(account, passphrase)
}

func getKeystore() *keystore.KeyStore {
	return keystore.NewKeyStore(config.GetKeystoreDir(), scryptN, scryptP)
}
