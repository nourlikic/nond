package args

import (
	"fmt"
	"github.com/nourlikic/nond/contracts"
	"github.com/nourlikic/nond/accounts"
	"github.com/nourlikic/nond/net"
)

func transaction(identifier string,method string){

	bytesJsonAccount,err := accounts.ReadEncDefaultAccount()
	if err != nil{
		fmt.Println(err)
		return
	}
	raw, err :=contracts.Transaction(identifier,method,bytesJsonAccount)
	if err != nil{
		fmt.Println(err)
		return
	}
	hash,err := net.SendRawTransaction(raw)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Transaction hash: " + hash)
}


func loadAbi(identifier string, path string) {
	abi, err := contracts.LoadAbi(identifier, path)
	if err != nil {
		fmt.Println(err)
	}
	for _, m := range abi.Methods {
		fmt.Println(m.String())
	}

}

func listAbi() {

	abis, err := contracts.ListAbis()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, abi := range abis {
		fmt.Println(abi)
	}
}

func detailAbi(identifier string) {

	abi, err := contracts.GetAbiObject(identifier)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, m := range abi.Methods {
		fmt.Println(m.String())
	}
}

func deleteAbi(identifier string) {

	err := contracts.DeleteAbi(identifier)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveContract(contractIdentifier string, address string, abiIdenfier string) {

	err := contracts.SaveContract(contractIdentifier, address, abiIdenfier)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func detailContract(identifier string) {

	address, abiIdentifier, err := contracts.GetContract(identifier)
	if err != nil {
		fmt.Println(err)
		return
	}
	abi, err := contracts.GetAbiObject(abiIdentifier)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Contract:  " + identifier)
	fmt.Println("Address:   " + address.Hex())
	fmt.Println("Abi:       " + abiIdentifier)
	for _, m := range abi.Methods {
		fmt.Println(m.String())
	}

}

func listContracts() {

	contracts, err := contracts.ListContracts()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, contract := range contracts {
		fmt.Println("+++")
		fmt.Println("Contract:  " + contract.ContractIdentifier)
		fmt.Println("Address:   " + contract.Address.Hex())
		fmt.Println("Abi:       " + contract.AbiIdentifier)
	}
	fmt.Println("+++")
}

func deleteContract(identifier string) {
	err := contracts.DeleteContract(identifier)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func changeContractAddress(identifier string, address string) {

}

func newAccount() {
	err := accounts.CreateNewAccount()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteAccount(address string, passphrase string) {

	err := accounts.DeleteAccount(address, passphrase)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func defaultAccount(address string) {

	err := accounts.DefaultAccount(address)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func listAccounts() {

	err := accounts.ListAccounts()
	if err != nil {
		fmt.Println(err)
		return
	}
}
