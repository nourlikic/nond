package contracts

import (
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/core/types"
	//"errors"
)



/*
func (c *Contract) CallMsg(
	opts *bind.CallOpts,
	result interface{},
	method string,
	params ...interface{}) (ethereum.CallMsg, error) {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	// Pack the input, call and unpack the results
	input, err := c.Abi.Pack(method, params...)
	if err != nil {
		return ethereum.CallMsg{}, err
	}
	return ethereum.CallMsg{From: opts.From, To: &c.Address, Data: input}, nil
}

func (c *Contract) TransactMsg(
	opts *bind.TransactOpts,
	method string,
	params ...interface{}) (*types.Transaction, error) {

	input, err := c.Abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	value := opts.Value
	nonce := opts.Nonce.Uint64()
	gasPrice := opts.GasPrice
	gasLimit := opts.GasLimit
	rawTx := types.NewTransaction(nonce, c.Address, value, gasLimit, gasPrice, input)
	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
*/


