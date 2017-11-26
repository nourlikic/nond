package main

import (
	"fmt"
	"os"
	"github.com/nourlikic/nond/args"
)

func main() {

	if !isInitialized() {
		fmt.Println("Will Initialize")
		initialize()
	}

	if len(os.Args) < 2 {
		fmt.Println("hi :)")
		os.Exit(0)
	}

	err := args.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}


/*
const (
	ActionRefresh               Action = "Refresh"
	ActionLoadAbi               Action = "LoadAbi"              // identity, path
	ActionListAbi               Action = "ListAbi"              // nil
	ActionDeleteAbi             Action = "DeleteAbi"            // identifier
	ActionDetailAbi             Action = "DetailAbi"            // identifier
	ActionSaveContract          Action = "SaveContract"         // address, abi
	ActionDeleteContract        Action = "DeleteContract"       // identifier
	ActionListContracts         Action = "ListContracts"        // nil
	ActionDetailContract        Action = "DetailContract"       // identifier
	ActionChangeContractAddress Action = "ChangeContractAddres" // identifier, address
	ActionNewAccount            Action = "NewAccount"           // identifier,passphrase
	ActionDeleteAccount         Action = "DeleteAccount"        // identifier
	ActionDefaultAccount        Action = "DefaultAccount"       //identifier
	ActionListAccounts          Action = "ListAccount"          // nil

	ActionCall           Action = "Call"        // cidentifier, args
	ActionTransaction    Action = "Transaction" // cidentifier, args
	ActionTestConnection Action = "TestConnection"
	ActionRelayUrl       Action = "SetRelayUrl"

	ActionChageGasPrice Action = "ChangeGasPrice"
	ActionChangeGas     Action = "ChangeGas"

	//Temp
	ActionIter Action = "iter"
)

type Action string
type Arg string

const (
	ArgContractIdentifier string = "cid"
	ArgAbiIdentifier      string = "aid"
	ArgPath               string = "path"
	ArgAddress            string = "address"
	ArgPassphrase         string = "passphrase"
)*/
