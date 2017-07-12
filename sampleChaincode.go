/*
Invoke Methods :
****************
register

Query Methods :
****************
GetTransactionInitDetailsForRefAndMaker

Dependency Methods :
*********************

GetTransactionInitiationMap
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Register struct {
}

type register struct {
	name   		string  `json:"name"`
   	uname   	string  `json:"uname"`
    emailid   	string  `json:"emailid"`
    password   	string  `json:"password"`
    usertype   	string  `json:"usertype"`
    acode   	string  `json:"acode"`
}

//Global declaration of maps
var register_map map[string]register

//Invoke methods starts here 

func registration(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var reg_obj register
	var err error

	fmt.Println("Entering register")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast one arguments for register")
	}

	//unmarshal register data from UI to "register" struct
	err = json.Unmarshal([]byte(args[1]), &reg_obj)
	if err != nil {
		fmt.Printf("Unable to unmarshal register input  : %s\n", err)
		return nil, nil
	}

	// saving register into map
	GetRegistrationMap(stub)	

	//put register data into map
	register_map[reg_obj.uname] = reg_obj	

	SetRegistrationMap(stub)	
	
	fmt.Printf("registration map : %v \n", register_map)	
	fmt.Println("Registration details Successfully saved")	
	
	return nil, nil
}


func GetRegistrationMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("RegMap")
	if err != nil {
		fmt.Printf("Failed to get RegMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("RegMap map exists.\n")
		err = json.Unmarshal(bytesread, &register_map)
		if err != nil {
			fmt.Printf("Failed to initialize RegMap for block chain :%v\n", err)
			return err
		}
	} else {
		fmt.Printf("RegMap map does not exist. To be created. \n")
		register_map = make(map[string]register)
		bytesread, err = json.Marshal(&register_map)
		if err != nil {
			fmt.Printf("Failed to initialize  RegMap for block chain :%v\n", err)
			return err
		}
		err = stub.PutState("RegMap", bytesread)
		if err != nil {
			fmt.Printf("Failed to initialize  RegMap for block chain :%v\n", err)
			return err
		}
	}
	return nil
}

//setTransactionInitiationMap
func SetRegistrationMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = json.Marshal(&register_map)
	if err != nil {
		fmt.Printf("Failed to set the RegMap for block chain :%v\n", err)
		return err
	}
	err = stub.PutState("RegMap", bytesread)
	if err != nil {
		fmt.Printf("Failed to set the RegMap %v\n", err)
		return errors.New("Failed to set the TransactionItemMap")
	}

	return nil
}

// Init sets up the chaincode
func (t *Register) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")
	return nil, nil
}

// Query the chaincode
func (t *Register) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//fmt.Println("Inside Query for test chaincode")
	//if function == "registration" {
	//	return registration(stub, args)
	//} 
	//return nil, nil
}

// Invoke the function in the chaincode
func (t *Register) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside invokve for test chaincode")
	if function == "registration" {
		return registration(stub,args)
	} 	
	fmt.Println("Function not found")
	return nil, nil
}

func main() {
	err := shim.Start(new(Register))
	if err != nil {
		fmt.Println("Could not start Register")
	} else {
		fmt.Println("Register successfully started")
	}

}


