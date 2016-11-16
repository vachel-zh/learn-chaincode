package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Test1Chaincode aaa
type Test1Chaincode struct {
}

func main() {
	err := shim.Start(new(Test1Chaincode))
	if err != nil {
		fmt.Printf("Error starting Test1 chaincode: %s", err)
	}
}

// Init resets all the things
func (t *Test1Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var userColDefs []*shim.ColumnDefinition
	userColDefs[0] = &shim.ColumnDefinition{"id", shim.ColumnDefinition_STRING, true}
	userColDefs[1] = &shim.ColumnDefinition{"phoneNo", shim.ColumnDefinition_STRING, false}
	err := stub.CreateTable("mg_user", userColDefs)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *Test1Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "insert" {
		var row shim.Row
		row.Columns[0] = &shim.Column{&shim.Column_String_{args[0]}}
		row.Columns[1] = &shim.Column{&shim.Column_String_{args[1]}}
		stub.InsertRow("mg_user", row)
		return nil, nil
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query aa
func (t *Test1Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		var key []shim.Column
		key[0] = shim.Column{&shim.Column_String_{args[0]}}
		row, err := stub.GetRow("mg_user", key)
		return []byte(row.String()), err
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
