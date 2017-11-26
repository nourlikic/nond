package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Contract struct {
	AbiIdentifier      string
	ContractIdentifier string
	Address            common.Address
	Abi                abi.ABI
}
