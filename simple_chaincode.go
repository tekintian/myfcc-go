package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"

)

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	
	if err != nil {
		fmt.Printf("Error Starting Simple Chaincode, Error: %s \n", err.Error())
	}
}
