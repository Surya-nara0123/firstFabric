package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("InitLedger")
	// Set initial value of the integer to 0
	err := ctx.GetStub().PutState("counter", []byte(strconv.Itoa(0)))
	if err != nil {
        return fmt.Errorf("failed to put to world state. %v", err)
    }
	return nil
}

// IncrementCounter increases the integer by 1
func (s *SmartContract) IncrementCounter(ctx contractapi.TransactionContextInterface) error {
	// Get the current value of the counter
	counterBytes, err := ctx.GetStub().GetState("counter")
	if err != nil {
		return fmt.Errorf("failed to get counter: %v", err)
	}
	if counterBytes == nil {
		return fmt.Errorf("counter not initialized")
	}

	// Convert the counter value to an integer
	counter, err := strconv.Atoi(string(counterBytes))
	if err != nil {
		return fmt.Errorf("failed to convert counter to integer: %v", err)
	}

	// Increment the counter
	counter++

	// Save the updated counter value back to the ledger
	return ctx.GetStub().PutState("counter", []byte(strconv.Itoa(counter)))
}

func (s *SmartContract) DecrementCounter(ctx contractapi.TransactionContextInterface) error {
	// Get the current value of the counter
	counterBytes, err := ctx.GetStub().GetState("counter")
	if err != nil {
		return fmt.Errorf("failed to get counter: %v", err)
	}
	if counterBytes == nil {
		return fmt.Errorf("counter not initialized")
	}

	// Convert the counter value to an integer
	counter, err := strconv.Atoi(string(counterBytes))
	if err != nil {
		return fmt.Errorf("failed to convert counter to integer: %v", err)
	}

	// Increment the counter
	counter--

	// Save the updated counter value back to the ledger
	return ctx.GetStub().PutState("counter", []byte(strconv.Itoa(counter)))
}

// GetCounter retrieves the current value of the counter
func (s *SmartContract) GetCounter(ctx contractapi.TransactionContextInterface) (int, error) {
	// Get the current value of the counter
	counterBytes, err := ctx.GetStub().GetState("counter")
	if err != nil {
		return 0, fmt.Errorf("failed to get counter: %v", err)
	}
	if counterBytes == nil {
		return 0, fmt.Errorf("counter not initialized")
	}

	// Convert the counter value to an integer and return it
	counter, err := strconv.Atoi(string(counterBytes))
	if err != nil {
		return 0, fmt.Errorf("failed to convert counter to integer: %v", err)
	}

	fmt.Println(counter)

	return counter, nil
}

func main() {
	cc1, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := cc1.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
