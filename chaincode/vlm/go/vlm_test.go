package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func printCar(t *testing.T, payload []byte) {
	var car CarStruct
	err := json.Unmarshal(payload, &car)
	if err != nil {
		t.Fatalf("MockInvoke error: Issue with Car json unmarshaling")
		return
	}
	fmt.Println("Car ->", car)
}

func testCreateCar(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testCreateCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("createCar"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)

}

func testTransferCar(t *testing.T, stub *shim.MockStub, chasisNo string, newOwner string) {
	fmt.Println("Entering testTransferCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("transferCar"), []byte(chasisNo), []byte(newOwner)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testSellnRegisterCar(t *testing.T,
	stub *shim.MockStub,
	chasisNo string,
	newOwner string,
	registrationNo string,
	registrationExp string,
	loanAmount int64) {
	fmt.Println("Entering testSellnRegisterCar")

	// Convert int64 to byte
	// bloanAmount := make([]byte, 8)
	// binary.LittleEndian.PutUint64(bloanAmount, uint64(loanAmount))

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("sellnRegisterCar"),
			[]byte(chasisNo),
			[]byte(newOwner),
			[]byte(registrationNo),
			[]byte(registrationExp),
			[]byte(strconv.FormatInt(loanAmount, 10))})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testScrapCar(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testScrapCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("scrapCar"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testRegisterClaim(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testRegisterClaim")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("registerClaim"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testClearLoan(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testClearLoan")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("clearLoan"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testIssueChallan(t *testing.T, stub *shim.MockStub, 
						chasisNo string,
						challanAmt int64) {
	fmt.Println("Entering testIssueChallan")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("issueChallan"), 
				[]byte(chasisNo),
				[]byte(strconv.FormatInt(challanAmt, 10))})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testPayChallan(t *testing.T, stub *shim.MockStub, 
	chasisNo string,
	challanAmt int64) {
	
	fmt.Println("Entering testPayChallan")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
	[][]byte{[]byte("payChallan"), 
	[]byte(chasisNo),
	[]byte(strconv.FormatInt(challanAmt, 10))})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testGetCar(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testGetCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("getCar"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	} else {
		printCar(t, result.Payload)
	}
}

func testGetCaByRegistrationNo(t *testing.T, stub *shim.MockStub, registrationNo string) {
	fmt.Println("Entering testGetCaByRegistrationNo")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("getCarByRegistrationNo"), []byte(registrationNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	} else {
		printCar(t, result.Payload)
	}
}

func testGetCarHistory(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testGetCarHistory")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("getCarHistory"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	fmt.Println(valAsbytes)
}

func TestSmartContract(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("vlm", scc)

	testCreateCar(t, stub, "1000")
	// testRegisterClaim(t, stub, "1000")
	testTransferCar(t, stub, "1000", "Rahul")
	testSellnRegisterCar(t, stub, "1000", "Vaibhav", "TS071", "31DEC2025", 202)
	// testGetCaByRegistrationNo(t, stub, "TS071")
	testRegisterClaim(t, stub, "1000")
	testRegisterClaim(t, stub, "1000")
	testIssueChallan(t, stub, "1000", 300)
	testIssueChallan(t, stub, "1000", 500)
	testClearLoan(t, stub, "1000")
	// testSellnRegisterCar(t, stub, "1000", "Vaibhav", "TS071", "31DEC2025", 202)
	testPayChallan(t, stub, "1000", 800)
	testClearLoan(t, stub, "1000")
	testScrapCar(t, stub, "1000")
	// testRegisterClaim(t, stub, "1000")
	// testSellnRegisterCar(t, stub, "1000", "Jyothi", "TS071", "31DEC2025")
	testGetCar(t, stub, "1000")
	// testGetCarHistory(t, stub, "1000")
}

