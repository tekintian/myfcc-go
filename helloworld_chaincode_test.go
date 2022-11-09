package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func checkInit(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shimtest.MockStub, name string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("get"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("query", name, "failed", string(res.Message))
		t.FailNow()
	}

	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	fmt.Println("Query value", name, " was ", string(res.Payload))
}

func checkInvoke(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", res.Message)
		t.FailNow()
	}
}

func Test_Helloworld(t *testing.T) {
	hello := new(HelloWorldChaincode)
	stub := shimtest.NewMockStub("hello", hello)

	checkInit(t, stub, [][]byte{[]byte("str"), []byte("helloworld!")})

	checkQuery(t, stub, "str")

	checkInvoke(t, stub, [][]byte{[]byte("set"), []byte("str"), []byte("hellowrold Tekin")})
	checkQuery(t, stub, "str")
}
