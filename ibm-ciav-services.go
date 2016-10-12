/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("Node-Web-Client-Test")

type ServicesChaincode struct {
}

func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("Hello, Init stared . . . ")
	err := stub.PutState("counter", []byte("0"))
	return nil, err
}


func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "abc" {
// 		stub.PutState("counter", counter)
		val1, err := stub.ReadCertAttribute("role")
		myLogger.Debug("Role : [%s]", val1)
		val2, err := stub.ReadCertAttribute("account")
		myLogger.Debug("Role : [%s]", account)
		
		stub.PutState("counter", []byte("1"))
	}
	return nil, nil
}

/*
 		Get Customer record by customer id or PAN number
*/
func (t *ServicesChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "read" {
		return read(stub, args)
	}
	return nil, errors.New("Invalid query function name. Expecting \"read\"")
}


func read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	// Get the state from the ledger
	Avalbytes, err := stub.GetState("counter")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for counter\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for counter\"}"
		return nil, errors.New(jsonResp)
	}

	isOk, _ := stub.VerifyAttribute("position123", []byte("Software Engineer"))
	var jsonResp string
	if isOk {
		jsonResp = "{\"counter1\" : " + string(Avalbytes) +
								"\"}"
	}else{
		jsonResp = "{\"counter2\" : " + string(Avalbytes) +
								"\"}"
	}

	fmt.Printf("Query Response:%s\n", jsonResp)

	bytes, err := json.Marshal(jsonResp)
	if err != nil {
		return nil, errors.New("Error converting kyc record")
	}
	return bytes, nil
}

func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}
