package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"errors"
	"github.com/nourlikic/nond/database"
	"fmt"
)

type Contract struct {
	address common.Address // Deployment address of the contract on the Ethereum blockchain
	abi     abi.ABI        // Reflect based ABI to access the correct Ethereum methods
}

func (c *Contract) CallMsg(
	opts *bind.CallOpts,
	result interface{},
	method string,
	params ...interface{}) (ethereum.CallMsg, error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	// Pack the input, call and unpack the results
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return ethereum.CallMsg{}, err
	}
	return ethereum.CallMsg{From: opts.From, To: &c.address, Data: input}, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (c *Contract) TransactMsg(
	opts *bind.TransactOpts,
	method string,
	params ...interface{}) (*types.Transaction, error) {

	//Input
	input, err := c.abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	//Value
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	//Nonce
	// retrieve nonce here
	nonce := opts.Nonce.Uint64()
	/*if nonce != {
		return nil,err
	}*/

	//GasPrice
	gasPrice := opts.GasPrice
	if gasPrice == nil {
		// get default?
	}

	gasLimit := opts.GasLimit
	if gasLimit == nil {
		// estimate gas or
	}

	rawTx := types.NewTransaction(nonce, c.address, value, gasLimit, gasPrice, input)

	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func ListContracts() error {
	contracts, err := database.ListContracts()
	if err != nil {
		return err
	}
	for _, contract := range contracts {
		fmt.Println("+++++++++")
		fmt.Println("contract: ", contract.ContractIdentifier)
		fmt.Println("address:  ", contract.Address)
		fmt.Println("abi:      ", contract.AbiIdentifier)
	}
	return nil
}

func SaveContract(contractIdentifier string, address string, abiIdentifier string) error {
	return database.SaveContract(contractIdentifier, address, abiIdentifier)
}

func DeleteContract(identifier string) error {
	return database.DeleteContract(identifier)
}

func DetailContract(identifier string) error {
	contract, err := database.DetailContract(identifier)
	if err != nil {
		return err
	}
	fmt.Println("contract: ", contract.ContractIdentifier)
	fmt.Println("address:  ", contract.Address)
	fmt.Println("abi:      ", contract.AbiIdentifier)
	return DetailAbi(contract.AbiIdentifier)

}
