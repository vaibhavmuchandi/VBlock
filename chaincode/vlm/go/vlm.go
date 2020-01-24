/*
 * This the smart contract for vehicle lifetime management
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// SmartContract structure
type SmartContract struct {
}

// CarStruct structure
type CarStruct struct {
	ChassisNo             string `json:"chassisNo"`
	Owner                 string `json:"owner"`
	RegistrationNo        string `json:"registrationNo"`
	RegistrationExpiryDae string `json:"registrationExpiryDae"`
	Status                string `json:"status"`
	LoanAmount            int64  `json:"loanAmt"`
	ChallanAmount         int64  `json:"challanAmt"`
	InsuranceClaim        int64  `json:"insuranceClaim"`
}

// Init SmartContract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke SmartContract Invoke
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "createCar" {
		return s.createCar(APIstub, args)
	} else if function == "transferCar" {
		return s.transferCar(APIstub, args)
	} else if function == "sellnRegisterCar" {
		return s.sellnRegisterCar(APIstub, args)
	} else if function == "scrapCar" {
		return s.scrapCar(APIstub, args)
	} else if function == "clearLoan" {
		return s.clearLoan(APIstub, args)
	} else if function == "issueChallan" {
		return s.issueChallan(APIstub, args)
	} else if function == "payChallan" {
		return s.payChallan(APIstub, args)
	} else if function == "registerClaim" {
		return s.registerClaim(APIstub, args)
	} else if function == "getCar" {
		return s.getCar(APIstub, args)
	} else if function == "getCarByRegistrationNo" {
		return s.getCarByRegistrationNo(APIstub, args)
	} else if function == "getCarHistory" {
		return s.getCarHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// createCar - This is for manufacture to create cars.
func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	chasisNo := args[0]

	Car := CarStruct{ChassisNo: chasisNo,
		Owner:                 "Maruti",
		RegistrationNo:        "",
		RegistrationExpiryDae: "",
		Status:                "New",
		LoanAmount:				0,
		ChallanAmount:			0,
		InsuranceClaim:			0}
	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("JSON Marshal failed.")
	}

	APIstub.PutState(chasisNo, CarBytes)
	fmt.Println("Car Created -> ", Car)

	return shim.Success(nil)
}

// transferCar - This is for manufacture to transfer the cars dealer.
func (s *SmartContract) transferCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	chasisNo := args[0]
	owner := args[1]

	CarAsBytes, _ := APIstub.GetState(chasisNo)

	var car CarStruct

	err := json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}

	if car.Status != "New" {
		return shim.Error("Only new car transfer is allowed")
	}

	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 owner,
		RegistrationNo:        "",
		RegistrationExpiryDae: "",
		Status:                "Dealer",
		LoanAmount:				0,
		ChallanAmount:			0,
		InsuranceClaim:			0}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car trasnferred to dealer -> ", Car)

	return shim.Success(nil)
}

func getCarForRegistrationNo(stub shim.ChaincodeStubInterface, registrationNo string) (string, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"registrationNo\":\"%s\"}}", registrationNo)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return "", err
		fmt.Println("Error in GetQueryResult")
	}
	defer resultsIterator.Close()

	fmt.Println("Has resulted ")
	for resultsIterator.HasNext() {
		fmt.Println("Iterator has cars")
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return "", err
		} else {
			return queryResponse.Key, nil
		}
	}
	return "", nil
}

// sellnRegisterCar - This is for dealers to sell the cars to customer.
func (s *SmartContract) sellnRegisterCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	chasisNo := args[0]
	owner := args[1]
	registrationNo := args[2]
	registrationExpiryDae := args[3]
	loanAmount, err := strconv.ParseInt(args[4], 10, 64)
	if err != nil {
		return shim.Error("Error converting loan amount.")
	}

	OldchassisNo, err := getCarForRegistrationNo(APIstub, registrationNo)
	if err == nil && OldchassisNo != "" && OldchassisNo != chasisNo {
		errStr := fmt.Sprintf("Car found for the same registration number. ChassisNo = %s", OldchassisNo)
		return shim.Error(errStr)
	}

	CarAsBytes, _ := APIstub.GetState(chasisNo)
	var car CarStruct

	err = json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}
	
	if car.LoanAmount != 0 {
		return shim.Error("Cannot sell car with loan.")
	}

	if car.ChallanAmount != 0 {
		return shim.Error("Cannot sell car with challan.")
	}

	if car.Status == "Scrapped" {
		return shim.Error("Cannot sell car that is already scrapped.")
	}

	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 owner,
		RegistrationNo:        registrationNo,
		RegistrationExpiryDae: registrationExpiryDae,
		Status:                "Customer",
		LoanAmount:				loanAmount,
		ChallanAmount:			0,
		InsuranceClaim:			car.InsuranceClaim}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car sold to customer -> ", Car)

	return shim.Success(nil)
}

// clearLoan - This is for bank to clear the car loans.
func (s *SmartContract) clearLoan(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	chasisNo := args[0]

	CarAsBytes, _ := APIstub.GetState(chasisNo)
	var car CarStruct

	err := json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}

	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 car.Owner,
		RegistrationNo:        car.RegistrationNo,
		RegistrationExpiryDae: car.RegistrationExpiryDae,
		Status:                car.Status,
		LoanAmount:				0,
		ChallanAmount:			car.ChallanAmount,
		InsuranceClaim:			car.InsuranceClaim}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car loan paid -> ", Car)

	return shim.Success(nil)
}

// issueChallan - This is for traffic police to issue challans for the car.
func (s *SmartContract) issueChallan(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	chasisNo := args[0]
	challanAmt, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return shim.Error("Error converting challan amount.")
	}

	CarAsBytes, _ := APIstub.GetState(chasisNo)
	var car CarStruct

	err = json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}

	newChallanAmt := challanAmt + car.ChallanAmount
	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 car.Owner,
		RegistrationNo:        car.RegistrationNo,
		RegistrationExpiryDae: car.RegistrationExpiryDae,
		Status:                car.Status,
		LoanAmount:				car.LoanAmount,
		ChallanAmount:			newChallanAmt,
		InsuranceClaim:			car.InsuranceClaim}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car challan issued -> ", Car)

	return shim.Success(nil)
}

// payChallan - This is for customer to pay challans for the car.
func (s *SmartContract) payChallan(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	chasisNo := args[0]
	challanAmt, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return shim.Error("Error converting challan amount.")
	}

	CarAsBytes, _ := APIstub.GetState(chasisNo)
	var car CarStruct

	err = json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}

	newChallanAmt := car.ChallanAmount - challanAmt;
	if newChallanAmt < 0 {
		newChallanAmt = 0
	}

	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 car.Owner,
		RegistrationNo:        car.RegistrationNo,
		RegistrationExpiryDae: car.RegistrationExpiryDae,
		Status:                car.Status,
		LoanAmount:				car.LoanAmount,
		ChallanAmount:			newChallanAmt,
		InsuranceClaim:			car.InsuranceClaim}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car challan -> ", Car)

	return shim.Success(nil)
}

// registerClaim - This is for insurer to register claim for the car.
func (s *SmartContract) registerClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	chasisNo := args[0]

	CarAsBytes, _ := APIstub.GetState(chasisNo)
	var car CarStruct

	err := json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}

	if car.Status != "Customer" {
		return shim.Error("Insurance claim can only be registered for customer car.")
	}

	newClainCount := car.InsuranceClaim + 1;

	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 car.Owner,
		RegistrationNo:        car.RegistrationNo,
		RegistrationExpiryDae: car.RegistrationExpiryDae,
		Status:                car.Status,
		LoanAmount:				car.LoanAmount,
		ChallanAmount:			car.ChallanAmount,
		InsuranceClaim:			newClainCount}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car insurance claim register -> ", Car)

	return shim.Success(nil)
}

// scrapCar - This is for customer to scarp the cars.
func (s *SmartContract) scrapCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	chasisNo := args[0]

	CarAsBytes, _ := APIstub.GetState(chasisNo)
	var car CarStruct

	err := json.Unmarshal(CarAsBytes, &car)
	if err != nil {
		return shim.Error("Issue with Car json unmarshaling")
	}

	if car.LoanAmount != 0 {
		return shim.Error("Cannot scrap car with loan.")
	}

	if car.ChallanAmount != 0 {
		return shim.Error("Cannot scarp car with challan.")
	}

	Car := CarStruct{ChassisNo: car.ChassisNo,
		Owner:                 car.Owner,
		RegistrationNo:        car.RegistrationNo,
		RegistrationExpiryDae: car.RegistrationExpiryDae,
		Status:                "Scrapped"}

	CarBytes, err := json.Marshal(Car)
	if err != nil {
		return shim.Error("Issue with Car json marshaling")
	}

	APIstub.PutState(Car.ChassisNo, CarBytes)
	fmt.Println("Car scrapped -> ", Car)

	return shim.Success(nil)
}

func (s *SmartContract) getCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	chasisNo := args[0]
	CarAsBytes, _ := APIstub.GetState(chasisNo)
	return shim.Success(CarAsBytes)
}

func (s *SmartContract) getCarByRegistrationNo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	registrationNo := args[0]
	chassisNo, err := getCarForRegistrationNo(APIstub, registrationNo)
	if err != nil {
		errStr := fmt.Sprintf("Car not found for the registration no. Error = %s", err)
		return shim.Error(errStr)
	}
	CarAsBytes, _ := APIstub.GetState(chassisNo)
	return shim.Success(CarAsBytes)
}

func (s *SmartContract) getCarHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	chasisNo := args[0]

	resultsIterator, err := APIstub.GetHistoryForKey(chasisNo)
	if err != nil {
		return shim.Error("Error retrieving Car history with GetHistoryForKey")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the car
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving next Car history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getCarHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// Main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
