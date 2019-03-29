/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Visa struct {
	Type            string `json:"type"`
	VisaCode        string `json:"visaCode"`
	PassNb          string `json:"passNb"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Autority        string `json:"autority"`
	DateOfExpiry    string `json:"dateOfExpiry"`
	DateOfIssue     string `json:"dateOfIssue"`
	PlaceOfIssue    string `json:"placeOfIssue"`
	Validity        string `json:"validity"`
	ValidFor        string `json:"validFor"`
	NumberOfEntries string `json:"numberOfEntries"`
	DurationOfStay  string `json:"durationOfStay "`
	Remarks         string `json:"remarks"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryVisa" {
		return s.queryVisa(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryVisa(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	visaAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(visaAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	visas := []Visa{
		Visa{Type: "P", VisaCode: "FR", PassNb: "14ML52147", Name: "Jean", Surname: "Dupont", Autority: "ddf", DateOfExpiry: "16/09/1985", DateOfIssue: "France", PlaceOfIssue: "Toulouse", Validity: "dfs", ValidFor: "Préfecture de ", NumberOfEntries: "Avenue des Facultés, 33400 Talence", DurationOfStay: "Marron", Remarks: "16/02/2023"},
		Visa{Type: "P", VisaCode: "FR", PassNb: "14ML52147", Name: "Jean", Surname: "Dupont", Autority: "ddf", DateOfExpiry: "16/09/1985", DateOfIssue: "France", PlaceOfIssue: "Toulouse", Validity: "dfs", ValidFor: "Préfecture de ", NumberOfEntries: "Avenue des Facultés, 33400 Talence", DurationOfStay: "Marron", Remarks: "16/02/2023"},
		Visa{Type: "P", VisaCode: "FR", PassNb: "14ML52147", Name: "Jean", Surname: "Dupont", Autority: "ddf", DateOfExpiry: "16/09/1985", DateOfIssue: "France", PlaceOfIssue: "Toulouse", Validity: "dfs", ValidFor: "Préfecture de ", NumberOfEntries: "Avenue des Facultés, 33400 Talence", DurationOfStay: "Marron", Remarks: "16/02/2023"},
	}

	i := 0
	for i < len(visas) {
		fmt.Println("i is ", i)
		visaAsBytes, _ := json.Marshal(visas[i])
		APIstub.PutState(strconv.Itoa(i), visaAsBytes)
		fmt.Println("Added", visas[i])
		i = i + 1
	}

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
