/*
Invoke Methods :
****************
register
Admin
user
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
	"time"
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
type stock struct {
	share_name  	string  `json:"share_name"`
   	price   		string  `json:"price"`
}

type accStock struct{
	uname   	string  `json:"uname"`
	share_list  []stock  `json:"share_list"`
}
//Global declaration of maps
var register_map map[string]register
var lsit_register_map map[string][]register
var stock_map map[string]stock
var list_stock_map map[string][]stock
var accStock_map map[string]accStock
var list_accStock_map map[string][]accStock

//Invoke methods starts here 

func register(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

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

func saveStocks(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var stock_obj stock
	var err error

	fmt.Println("Entering saveStocks")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast one arguments for register")
	}

	//unmarshal stock data from UI to "register" struct
	err = json.Unmarshal([]byte(args[1]), &stock_obj)
	if err != nil {
		fmt.Printf("Unable to unmarshal register input  : %s\n", err)
		return nil, nil
	}

	// saving stock into map
	GetStockMap(stub)	

	//put stock data into map
	stock_map[stock_obj.share_name] = stock_obj	

	SetStockMap(stub)	
	
	fmt.Printf("stock map : %v \n", stock_map)	
	fmt.Println("stock details Successfully saved")	
	
	return nil, nil
}

//GetRegistrationMap
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

//SetRegistrationMap
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

//getStockMap
func GetStockMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("StockMap")
	if err != nil {
		fmt.Printf("Failed to get RegMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("StockMap map exists.\n")
		err = json.Unmarshal(bytesread, &register_map)
		if err != nil {
			fmt.Printf("Failed to initialize StockMap for block chain :%v\n", err)
			return err
		}
	} else {
		fmt.Printf("StockMap map does not exist. To be created. \n")
		register_map = make(map[string]register)
		bytesread, err = json.Marshal(&register_map)
		if err != nil {
			fmt.Printf("Failed to initialize  StockMap for block chain :%v\n", err)
			return err
		}
		err = stub.PutState("StockMap", bytesread)
		if err != nil {
			fmt.Printf("Failed to initialize  StockMap for block chain :%v\n", err)
			return err
		}
	}
	return nil
}

//setStockMap
func SetStockMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = json.Marshal(&register_map)
	if err != nil {
		fmt.Printf("Failed to set the StockMap for block chain :%v\n", err)
		return err
	}
	err = stub.PutState("StockMap", bytesread)
	if err != nil {
		fmt.Printf("Failed to set the StockMap %v\n", err)
		return errors.New("Failed to set the TransactionItemMap")
	}

	return nil
}

func validateUsername(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var res="NOTEXIST";

	fmt.Println("Entering validateUsername")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing username")
	}
	fmt.Printf("Entering validateUsername : %v\n", args[0])
	var username = args[0]
	GetRegistrationMap();
	for _, value := range register_map {
		if value.uname == username {
			res="EXIST";
			break;
		}
	}

	fmt.Printf("username : %v\n", res)
	return res,nil;
}

func validateLogin(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	var res="NOTEXIST";

	fmt.Println("Entering validateLogin")

	if len(args) < 3 
	{
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing username or password or usertype")
	}

	fmt.Printf("Entering validateLogin username : %v\n", args[0])
	fmt.Printf("Entering validateLogin password : %v\n", args[1])
	fmt.Printf("Entering validateLogin usertype : %v\n", args[2])

	var username = args[0]
	var password = args[1]
	var usertype = args[2]

	GetRegistrationMap();

	for _, value := range register_map {
		if value.uname == username {
			if value.password == password {
				if value.usertype == usertype
				{
				res="EXIST";
				break;
				}
			}			
		}
	}

	fmt.Printf("username and password : %v\n", res)
	return res,nil;
}


func getAccountList(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("RegMap")
	if err != nil {
		fmt.Printf("Failed to get the RegMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("RegMap map exists.\n")

		list_register_map = make(map[string][]register)
		err = json.Unmarshal(bytesRead, &list_register_map)
		if err != nil {
			fmt.Printf("Failed to initialize the RegMap for block chain :%v\n", err)
			return err
		}
	}
	return bytesRead,nil
}


func getStockList(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("RegMap")
	if err != nil {
		fmt.Printf("Failed to get the RegMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("RegMap map exists.\n")

		list_stock_map = make(map[string][]stock)
		err = json.Unmarshal(bytesRead, &list_stock_map)
		if err != nil {
			fmt.Printf("Failed to initialize the RegMap for block chain :%v\n", err)
			return err
		}
	}
	return bytesRead,nil
}

func UpdatePriceofStock(stub shim.ChaincodeStubInterface, args []string) error {

	var stock_obj stock
	var ok bool
	fmt.Println("Entering Updatestock")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return errors.New("Missing stock")
	}

	var share_name=args[1];
	var price=args[2];
	fmt.Printf("share_name : %v\n", share_name+" ; price :  %v\n"+price)

	getStockList(stub)
	fmt.Printf("stock map : %v\n", list_stock_map)

	stock_obj.share_name=share_name
	stock_obj.price=price
	list_stock_map[share_name]=stock_obj;

	setStockList(stub)

	fmt.Printf("stock map : %v\n", list_stock_map)
	fmt.Println("Successfully saved item")
	return nil
}

func updateAccountStock(stub shim.ChaincodeStubInterface, args []string) error {

	var accStock_obj accStock_map
	var stock_array []stock
	var loc_stock stock
	var ok bool
	fmt.Println("Entering updateAccountStock")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return errors.New("Missing stock")
	}

    // put arguments into stock struct
	var uname=args[1]
	var share_name=args[2]
	var price=args[3]

	fmt.Printf("uname :  %v\n"+uname+"share_name : %v\n", share_name+" ; price :  %v\n"+price)

	loc_stock.share_name=share_name
	loc_stock.price=price  
    
	// get existing stock list from blockchain

	

	getAccStockList(stub)
	fmt.Printf("AccStock map : %v\n", list_accStock_map)

	accStock_obj=list_accStock_map[uname]
	stock_array=accStock_obj.share_list
	stock_array.add
	list_accStock_map[uname]=accStock_obj

	setAccStockList(stub)

	fmt.Printf("AccStock map : %v\n", list_accStock_map)
	fmt.Println("Successfully saved item")
	return nil
}


// Init sets up the chaincode
func (t *SBITransaction) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")
	return nil, nil
}

// Query the chaincode
func (t *SBITransaction) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	if function == "validateUsername" {
		return validateUsername(stub, args)
	} else if function == "validateLogin" {
		return validateLogin(stub, args)
	} else if function == "getAccountList" {
		return getAccountList(stub, args)
	} else if function == "getStockList" {
		return getStockList(stub, args)
	} 
	
	return nil, nil
}

// Invoke the function in the chaincode
func (t *Register) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "register" {
		return register(stub,args)
	} else if function == "UpdatePriceofStock" {
		return UpdatePriceofStock(stub,args)
	} 	
	
	fmt.Println("Function not found")
	return nil, nil
}

func main() {
	err := shim.Start(new(Register))
	if err != nil {
		fmt.Println("Could not start Registration")
	} else {
		fmt.Println("Registration successfully started")
	}

}



