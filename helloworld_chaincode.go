package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)
// 定义一个接口
type HelloWorldChaincode struct {
}
// 初始化方法
func (t *HelloWorldChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取字符串参数
	args:=stub.GetStringArgs()

	err:=stub.PutState(args[0],[]byte(args[1]))

	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success(nil)
}
// 引用方法
func (t *HelloWorldChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args:=stub.GetFunctionAndParameters()

	if fn=="set" {
		return t.set(stub, args)
	}else if fn=="get" {
		return t.get(stub,args)
	}

	return shim.Success(nil)
}
func (t *HelloWorldChaincode) set(stub shim.ChaincodeStubInterface, args []string) peer.Response  {
	err:=stub.PutState(args[0],[]byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
func (t *HelloWorldChaincode) get(stub shim.ChaincodeStubInterface, args []string) peer.Response  {
	val,err:=stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(val)
}
func main() {
	err := shim.Start(new(HelloWorldChaincode))

	if err != nil {
		fmt.Printf("Error Starting Simple Chaincode, Error: %s \n", err.Error())
	}
}
