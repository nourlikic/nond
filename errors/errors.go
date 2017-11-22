package errors

import "errors"

var (

	ErrorNoSuchFile = errors.New("No Such File")
	ErrorShowUsage = errors.New("Show Usage")
	ErrorNotExistAbi = errors.New("Abi Not Exist")
	ErrorAlreadyExistAbi = errors.New("Abi Identifier is Already Exist")
	ErrorNotExistAccount = errors.New("Account Not Exist")
	ErrorNotExistContract = errors.New("Contract Not Exist")
	ErrorAlreadyExistContract = errors.New("Contract Identifier is Already Exist")
)
