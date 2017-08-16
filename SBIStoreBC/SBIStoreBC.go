
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
    sbiStruct "github.com/RathikaSnowflake/SBIStruct1/sbiStruct"
)

// transaction will implement the processes
type SBITransaction struct {
}


//Global declaration of maps
var user_map map[string]sbiStruct.user

//Invoke methods starts here 

func registerUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var user_obj sbiStruct.user	
	var err error

	fmt.Println("Entering registerUser")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast one arguments for initiate Transaction")
	}

	fmt.Println("Args [0] is : %v\n",args[0])
	fmt.Println("Args [1] is : %v\n",args[1])
	
	//unmarshal transaction initiation data from UI to "transactionInitiation" struct
	err = json.Unmarshal([]byte(args[1]), &user_obj)
	if err != nil {
		fmt.Printf("Unable to unmarshal createTransaction input transaction initiation : %s\n", err)
		return nil, nil
	}

	fmt.Println("TransactionInitiation object refno variable value is : %s\n",trans_obj.TransRefNo);
	
	GetUserMap(stub)	

	user_map[user_obj.uname] = user_obj	

	SetUserMap(stub)	
	
	fmt.Printf("final user map : %v \n", user_map)		
	
	return nil, nil
}

func GetUserMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("userMap")
	if err != nil {
		fmt.Printf("Failed to get  Transaction initiation for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("userMap map exists.\n")
		err = json.Unmarshal(bytesread, &user_map)
		if err != nil {
			fmt.Printf("Failed to initialize  userMap for block chain :%v\n", err)
			return err
		}
	} else {
		fmt.Printf("userMap map does not exist. To be created. \n")
		user_map = make(map[string]sbiStruct.user	)
		bytesread, err = json.Marshal(&user_map)
		if err != nil {
			fmt.Printf("Failed to initialize  userMap for block chain :%v\n", err)
			return err
		}
		err = stub.PutState("userMap", bytesread)
		if err != nil {
			fmt.Printf("Failed to initialize  userMap for block chain :%v\n", err)
			return err
		}
	}
	return nil
}


//setTransactionInitiationMap
func SetUserMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = json.Marshal(&user_map)
	if err != nil {
		fmt.Printf("Failed to set the userMap for block chain :%v\n", err)
		return err
	}
	err = stub.PutState("userMap", bytesread)
	if err != nil {
		fmt.Printf("Failed to set the userMap %v\n", err)
		return errors.New("Failed to set the userMap")
	}

	return nil
}


// Init sets up the chaincode
func (t *SBITransaction) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")
	return nil, nil
}

// Query the chaincode
func (t *SBITransaction) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) { 	
	return nil, nil
}

// Invoke the function in the chaincode
func (t *SBITransaction) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "registerUser" {
		return registerUser(stub,args)
	}  
	fmt.Println("Function not found")
	return nil, nil
}

func main() {
	err := shim.Start(new(SBITransaction))
	if err != nil {
		fmt.Println("Could not start SBITransaction")
	} else {
		fmt.Println("SBITransaction successfully started")
	}

}


