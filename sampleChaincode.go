/* Created by Rathika Muthuraman

Invoke Methods :
****************
RegisterUser
SaveStocks
UpdateAccountStock
UpdateAccountBalance
UpdatePriceofStock
BuySellStock

Query Methods :
****************

ValidateUsername
ValidateLogin
GetAccountList
GetStockList
ListStockForAcc
GetAccountDetails

Dependency methods :
********************


*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type DVP_empty struct {
}


type user struct {
	Name   		string  `json:"name"`
   	Uname   	string  `json:"uname"`
    Emailid   	string  `json:"emailid"`
    Password   	string  `json:"password"`
    Usertype   	string  `json:"usertype"`
    Acode   	string  `json:"acode"`
    CashBalance string  `json:"cashBalance"`
}
type stock struct {
	Share_name  	string  `json:"share_name"`
   	Price   		string  `json:"price"`
    Qty             string  `json:"quantity"`
}

type accStock struct{
	Uname   	string  `json:"uname"`
	Share_list  []stock  `json:"share_list"`   
}

//Global declaration of maps
var user_map map[string]user
var stock_map map[string]stock
var accStock_map map[string]accStock

//Invoke methods starts here 

func RegisterUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var regObj user
	var err error

	fmt.Println("Entering registerUser ")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast one arguments for registering User")
	}
	fmt.Println("args : "+args[1])
	//unmarshal user data from UI to "user" struct
	err = json.Unmarshal([]byte(args[1]), &regObj)
	fmt.Printf("regObj : %v \n",regObj.Uname)	
	if err != nil {
		fmt.Printf("Unable to unmarshal user input  : %s\n", err)
		return nil, nil
	}

	// getting usermap
	GetUserMap(stub)	
	fmt.Printf("regObj 2 : %v \n",regObj.Uname)	
	//put user data into map
	user_map[regObj.Uname] = regObj	
	fmt.Printf("user map 1 : %v \n",user_map)	
	// saving user into map
	SetUserMap(stub)	
	
	fmt.Printf("user map : %v \n",user_map)	
	fmt.Println("Registration details Successfully saved")	
	
	return nil, nil
}

func SaveStocks(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var stock_obj stock
	var err error

	fmt.Println("Entering saveStocks")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast one arguments for user")
	}

	//unmarshal stock data from UI to "user" struct
	err = json.Unmarshal([]byte(args[1]), &stock_obj)
	if err != nil {
		fmt.Printf("Unable to unmarshal user input  : %s\n", err)
		return nil, nil
	}

	// saving stock into map
	GetStockMap(stub)	

	//put stock data into map
	stock_map[stock_obj.Share_name] = stock_obj	

	SetStockMap(stub)	
	
	fmt.Printf("stock map : %v \n", stock_map)	
	fmt.Println("stock details Successfully saved")	
	
	return nil, nil
}

//GetRegistrationMap
func GetUserMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	fmt.Printf("Entered into getUserMap")
	bytesread, err = stub.GetState("UserMap")
	if err != nil {
		fmt.Printf("Failed to get UserMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("UserMap map exists.\n")
		err = json.Unmarshal(bytesread, &user_map)
		if err != nil {
			fmt.Printf("Failed to initialize UserMap for block chain :%v\n", err)
			return err
		}
	} else {
		fmt.Printf("UserMap map does not exist. To be created. \n")
		user_map = make(map[string]user)
		bytesread, err = json.Marshal(&user_map)
		if err != nil {
			fmt.Printf("Failed to initialize  UserMap for block chain :%v\n", err)
			return err
		}
		err = stub.PutState("UserMap", bytesread)
		if err != nil {
			fmt.Printf("Failed to initialize  UserMap for block chain :%v\n", err)
			return err
		}
	}
	return nil
}

//SetRegistrationMap
func SetUserMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	fmt.Printf("Entered into setUserMap")
	bytesread, err = json.Marshal(&user_map)
	if err != nil {
		fmt.Printf("Failed to set the UserMap for block chain :%v\n", err)
		return err
	}
	err = stub.PutState("UserMap", bytesread)
	if err != nil {
		fmt.Printf("Failed to set the UserMap %v\n", err)
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
		fmt.Printf("Failed to get UserMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("StockMap map exists.\n")
		err = json.Unmarshal(bytesread, &stock_map)
		if err != nil {
			fmt.Printf("Failed to initialize StockMap for block chain :%v\n", err)
			return err
		}
	} else {
		fmt.Printf("StockMap map does not exist. To be created. \n")
		stock_map = make(map[string]stock)
		bytesread, err = json.Marshal(&stock_map)
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

	bytesread, err = json.Marshal(&stock_map)
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

//getAccStockMap
func GetAccStockMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("AccStockMap")
	if err != nil {
		fmt.Printf("Failed to get AccStockMap for block chain :%v\n", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("StockMap map exists.\n")
		err = json.Unmarshal(bytesread, &accStock_map)
		if err != nil {
			fmt.Printf("Failed to initialize AccStockMap for block chain :%v\n", err)
			return err
		}
	} else {
		fmt.Printf("AccStockMap map does not exist. To be created. \n")
		accStock_map = make(map[string]accStock)
		bytesread, err = json.Marshal(&accStock_map)
		if err != nil {
			fmt.Printf("Failed to initialize  AccStockMap for block chain :%v\n", err)
			return err
		}
		err = stub.PutState("AccStockMap", bytesread)
		if err != nil {
			fmt.Printf("Failed to initialize  AccStockMap for block chain :%v\n", err)
			return err
		}
	}
	return nil
}

//setAccStockMap
func SetAccStockMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = json.Marshal(&accStock_map)
	if err != nil {
		fmt.Printf("Failed to set the AccStockMap for block chain :%v\n", err)
		return err
	}
	err = stub.PutState("AccStockMap", bytesread)
	if err != nil {
		fmt.Printf("Failed to set the AccStockMap %v\n", err)
		return errors.New("Failed to set the TransactionItemMap")
	}

	return nil
}


func ValidateLogin(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var res="NOTEXIST";
	var bytesread []byte
	var err error
	
	fmt.Println("Entering validateLogin")

	if (len(args) < 3 )	{
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing username or password or usertype")
	}

	fmt.Printf("Entering validateLogin username : %v\n", args[1])
	fmt.Printf("Entering validateLogin password : %v\n", args[2])
	fmt.Printf("Entering validateLogin usertype : %v\n", args[3])

	var username = args[1]
	var password = args[2]
	var usertype = args[3]

	GetUserMap(stub);

	for _, value := range user_map {
		if value.Uname == username {
			if value.Password == password {
				if value.Usertype == usertype{
				res="EXIST";
				break;
				}
			}			
		}
	}
	
	bytesread,err=json.Marshal(&res)

	fmt.Printf("username and password : %v\n", res)
	return bytesread,err;
}


func GetAccountList(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var err error
	var bytesread []byte
    var list_user_map []user

	bytesread, err = stub.GetState("UserMap")
	if err != nil {
		fmt.Printf("Failed to get the UserMap for block chain :%v\n", err)
		return nil,err
	}

	if len(bytesread) != 0     {
		fmt.Printf("UserMap map exists.\n")
        //list_user_map = make(map[string][]user)
	    for _, value := range user_map {
		    list_user_map = append(list_user_map, value)
	    }
	    fmt.Printf("list of AllTransactions : %v\n", list_user_map)

	    bytesread, err = json.Marshal(&list_user_map)

		if err != nil {
			fmt.Printf("Failed to initialize the UserMap for block chain :%v\n", err)
			return nil,err
		}
	}
	return bytesread,nil
}




func GetStockList(stub shim.ChaincodeStubInterface) ([]byte, error) {
	var err error
	var bytesread []byte
    var list_stock_map []stock

	bytesread, err = stub.GetState("UserMap")
	if err != nil {
		fmt.Printf("Failed to get the UserMap for block chain :%v\n", err)
		return nil,err
	}
	if len(bytesread) != 0 {
		fmt.Printf("UserMap map exists.\n")

        //list_stock_map = make(map[string][]stock)
	    for _, value := range stock_map {
		    list_stock_map = append(list_stock_map, value)
	    }
	    fmt.Printf("list of AllTransactions : %v\n", list_stock_map)

	    bytesread, err = json.Marshal(&list_stock_map)

		if err != nil {
			fmt.Printf("Failed to initialize the UserMap for block chain :%v\n", err)
			return nil,err
		}
	}
	return bytesread,nil
}

func GetAccountDetails(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
	var err error
	var bytesread []byte
    var user_obj user

    fmt.Printf("Entering validateUsername : %v\n", args[1])

    var uname=args[1]
    GetUserMap(stub);

    user_obj=user_map[uname]

	bytesread, err = json.Marshal(&user_obj)

	if err != nil {
		fmt.Printf("Unable to marshal user_obj  %s\n", err)
		return nil, err
	}

	fmt.Printf(" user_obj : %v\n", bytesread)

	return bytesread,nil
}

func ListStockForAcc(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
	var err error
	var bytesread []byte
    var accStock_obj accStock

    if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing username")
	}
	fmt.Printf("Entering validateUsername : %v\n", args[1])

    var uname=args[1]
    GetAccStockMap(stub);

    accStock_obj=accStock_map[uname]

	bytesread, err = json.Marshal(&accStock_obj)

	if err != nil {
		fmt.Printf("Unable to marshal accStock_obj  %s\n", err)
		return nil, err
	}

	fmt.Printf(" Stocks for particular account : %v\n", bytesread)

	return bytesread,nil
}

func ValidateUsername(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var res="NOTEXIST";
	var bytesread []byte
	var err error
	fmt.Println("Entering validateUsername")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing username")
	}
	fmt.Printf("Entering validateUsername : %v\n", args[1])
	var username = args[1]
	GetUserMap(stub);
	for _, value := range user_map {
		if value.Uname == username {
			res="EXIST";
			break;
		}
	}
	bytesread,err=json.Marshal(&res)
	fmt.Printf("username : %v\n", res)
	return bytesread,err;
}
func UpdatePriceofStock(stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {

	var stock_obj stock
	
	fmt.Println("Entering Updatestock")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil,errors.New("Missing stock")
	}

	var share_name=args[1];
	var price=args[2];
	fmt.Printf("share_name : %v\n", share_name+" ; price :  %v\n"+price)

	GetStockMap(stub)
	fmt.Printf("stock map : %v\n", stock_map)

	stock_obj.Share_name=share_name
	stock_obj.Price=price
	stock_map[share_name]=stock_obj;

	SetStockMap(stub)

	fmt.Printf("stock map : %v\n", stock_map)
	fmt.Println("Successfully saved the stock")
	return nil,nil
}

func UpdateAccountStock(stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {

	var accStock_obj accStock
	var stock_array []stock
	var stock_obj stock

	
	fmt.Println("Entering into method updateAccountStock")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil,errors.New("Missing stock")
	}

    // put arguments into stock struct
	var uname=args[1]
	fmt.Println(" arguments "+uname)
	var share_name=args[2]
	var price=args[3]
    var qty=args[4]

	fmt.Printf("uname :  %v\n"+uname+"share_name : %v\n", share_name+" ; price :  %v\n"+price)

	stock_obj.Share_name=share_name
	stock_obj.Price=price  
    stock_obj.Qty=qty

	GetAccStockMap(stub)                   // get existing stock list from blockchain
	fmt.Printf("AccStock map : %v\n", accStock_map )

	accStock_obj=accStock_map [uname]       // fetch accstock for this particular userid

	stock_array=accStock_obj.Share_list     // fetch existing share list for this userid.need to check
	//stock_array.add(stock_obj)              // add new stock with existing stock array for this user.need to check
    stock_array=append(stock_array, stock_obj)
    accStock_obj.Share_list=stock_array     // put the newly created stock array in the struct object. need to check
	accStock_map[uname]=accStock_obj        // update newly updated object into map. need to check

	SetAccStockMap(stub)                   // update newly created map into blockchain

	fmt.Printf("AccStock map : %v\n", accStock_map )
	fmt.Println("Stock saved Successfully")
	return nil,nil
}


func UpdateAccountBalance(stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {		
			
        var user_obj user

		fmt.Printf("updateAccountBalance\n")

		if 	(len(args) < 2) {
		fmt.Println("Invalid number of args")
		return nil,errors.New("Expected atleast one arguments for updateAccountBalance" + args[0])
		}

	    var uname=args[1]
	    var cashBalance=args[2]

		fmt.Printf("uname :  %v\n"+uname+"cashBalance : %v\n", cashBalance)

        // get user values from blockchain
		GetUserMap(stub)

		user_obj=user_map[uname]

		//update cashBalance
		user_obj.CashBalance=cashBalance

		//update new struct values
		user_map[uname]=user_obj

		//update new values in blockchain
		SetUserMap(stub)	

		return nil,nil	
}

func BuySellStock(stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {	
	fmt.Println("Entering buySellStock")

	if len(args) < 8 {
		fmt.Println("Invalid number of arguments")
		return nil,errors.New("Missing stock")
	}
 
	var b_uname=args[1]
	var b_share_name=args[2]
	var b_qty=args[3]
    var b_cashBalance=args[4]	

	var s_uname=args[5]
	var s_share_name=args[6]
	var s_qty=args[7]
    var s_cashBalance=args[8]

	fmt.Printf("b_uname :  %v\n"+b_uname+"b_share_name : %v\n", b_share_name+" ; b_qty :  %v\n"+b_qty+" ; b_cashBalance :  %v\n"+b_cashBalance)
	fmt.Printf("s_uname :  %v\n"+s_uname+"s_share_name : %v\n", s_share_name+" ; s_qty :  %v\n"+s_qty+" ; s_cashBalance :  %v\n"+s_cashBalance)

   // put arguments into stock struct
    BuySellStockUpdate(stub,b_uname,b_share_name,b_qty,b_cashBalance,"buyer")
    BuySellStockUpdate(stub,s_uname,s_share_name,s_qty,s_cashBalance,"seller")

	return nil,nil
}

func BuySellStockUpdate(stub shim.ChaincodeStubInterface, uname string,share_name string,qty string,cashBalance string,user_buysell string) ([]byte,error) {

  	var stock_array []stock
	var stock_obj stock
	var accStock_obj accStock
	var user_obj user 	

	GetAccStockMap(stub)                   // get existing stock list from blockchain
	fmt.Printf("AccStock map : %v\n", accStock_map )

	accStock_obj=accStock_map [uname]       // fetch accstock for this particular userid
	stock_array=accStock_obj.Share_list     // fetch existing share list for this userid.need to check

    var length=len(stock_array)
    var exist="NOT EXIST"
    var j=1
    if length>0 {
	for i := 0; i < length; i++ {
        stock_obj=stock_array[i]
        if stock_obj.Share_name == share_name {
			stock_obj.Share_name=share_name

			if(user_buysell=="buyer"){
					stock_obj.Qty=stock_obj.Qty+qty
			} else if(user_buysell=="seller"){
					stock_obj.Qty=qty
			} 

            stock_obj.Qty=qty

            stock_array[i]=stock_obj
            exist="EXIST"
			break;
			}
		
		fmt.Println(i)
		j++;
		}
	if exist=="NOT EXIST" {
		stock_obj.Share_name=share_name
        stock_obj.Qty=qty
 		stock_array[j]=stock_obj
	}
	} else{
		stock_obj.Share_name=share_name
        stock_obj.Qty=qty
		stock_array[0]=stock_obj
	}
	
    // Update accStock struct
    accStock_obj.Share_list=stock_array     // put the newly created stock array in the struct object. need to check
	accStock_map[uname]=accStock_obj        // update newly updated object into map. need to check

	SetAccStockMap(stub)                   // update newly created map into blockchain

	fmt.Printf("AccStock map : %v\n", accStock_map )
	fmt.Println("Stock saved Successfully")

        // Update user sturct
	    GetUserMap(stub)  
        fmt.Printf("user map : %v\n", user_map ) 
        user_obj=user_map[uname]

		if(user_buysell=="seller"){
				user_obj.CashBalance=user_obj.CashBalance+cashBalance
			}	else if(user_buysell=="buyer"){
				user_obj.CashBalance=cashBalance
			}

        SetUserMap(stub) 

		return nil,nil

}

// Init sets up the chaincode
func (t *DVP_empty) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Inside INIT for test chaincode")
	return nil, nil
}

// Query the chaincode
func (t *DVP_empty) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {	
	if function == "validateUsername" {
		return ValidateUsername(stub, args)
	} else if function == "ValidateLogin" {
		return ValidateLogin(stub, args)
	} else if function == "GetAccountList" {
		return GetAccountList(stub)
	} else if function == "GetStockList" {
		return GetStockList(stub)
	} else if function == "ListStockForAcc" {
		return ListStockForAcc(stub, args)
	} else if function == "GetAccountDetails" {
		return GetAccountDetails(stub, args)
	}  

	return nil, nil
}

// Invoke the function in the chaincode
func (t *DVP_empty) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "RegisterUser" {
		return RegisterUser(stub,args)
	} else if function == "UpdateAccountStock" {
		return UpdateAccountStock(stub,args)
	} else if function == "UpdateAccountBalance" {
		return UpdateAccountBalance(stub,args)
	} else if function == "UpdatePriceofStock" {
		return UpdatePriceofStock(stub,args)
	} else if function == "SaveStocks" {
		return SaveStocks(stub,args)
	} else if function == "BuySellStock" {
		return BuySellStock(stub,args)
	}	
	
	
	
	fmt.Println("Function not found")
	return nil, nil
}

func main() {
	err := shim.Start(new(DVP_empty))
	if err != nil {
		fmt.Println("Could not start Registration")
	} else {
		fmt.Println("Registration successfully started")
	}

}


