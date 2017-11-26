package net

import "math/big"

func GetNonce() *big.Int{
	return big.NewInt(0)
}


func SendRawTransaction(rawData []byte) (hash string,err error){

	return "",nil
}