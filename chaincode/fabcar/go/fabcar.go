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
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the passport structure, with 4 properties.  Structure tags are used by encoding/json library
type Passport struct {
	Type         string  `json:"type"`
	CountryCode  string  `json:"countryCode"`
	PassNb       string  `json:"passNb"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	DateOfBirth  string  `json:"dateOfBirth"`
	Nationality  string  `json:"nationality"`
	Sex          string  `json:"sex"`
	PlaceOfBirth string  `json:"placeOfBirth"`
	Height       float64 `json:"height"`
	Autority     string  `json:"autority"`
	Residence    string  `json:"residence"`
	EyesColor    string  `json:"eyesColor"`
	DateOfExpiry string  `json:"dateOfExpiry"`
	DateOfIssue  string  `json:"dateOfIssue"`
	PassOrigin   string  `json:"passOrigin"`
	Validity     string  `json:"validity"`
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
	if function == "queryPassport" {
		return s.queryPassport(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createPassport" {
		return s.createPassport(APIstub, args)
	} else if function == "queryAllPassports" {
		return s.queryAllPassports(APIstub)
	} else if function == "changePassportOwner" {
		return s.changePassportOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryPassport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	passportAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(passportAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	taille := 1.65
	passports := []Passport{
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML52147", Name: "Jean", Surname: "Dupont", DateOfBirth: "10/05/1995", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "25/01/2015", PassOrigin: "France", Validity: "Valide"},
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML52147", Name: "Brad", Surname: "Dupont", DateOfBirth: "10/05/1995", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "25/01/2015", PassOrigin: "France", Validity: "Valide"},
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML52147", Name: "Jin Soo", Surname: "Dupont", DateOfBirth: "10/05/1995", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "25/01/2015", PassOrigin: "France", Validity: "Valide"},
	}

	i := 0
	for i < len(passports) {
		fmt.Println("i is ", i)
		passportAsBytes, _ := json.Marshal(passports[i])
		APIstub.PutState(strconv.Itoa(i), passportAsBytes)
		fmt.Println("Added", passports[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createPassport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	taille, _ := strconv.ParseFloat(args[10], 64)

	var passport = Passport{Type: args[1], CountryCode: args[2], PassNb: args[3], Name: args[4], Surname: args[5], DateOfBirth: args[6], Nationality: args[7], Sex: args[8], PlaceOfBirth: args[9], Height: taille, Autority: args[11], Residence: args[12], EyesColor: args[13], DateOfExpiry: args[14], DateOfIssue: args[15], PassOrigin: args[16], Validity: args[17]}

	passportAsBytes, _ := json.Marshal(passport)
	APIstub.PutState(args[0], passportAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllPassports(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"id\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"infos\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllPassports:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changePassportOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	passportAsBytes, _ := APIstub.GetState(args[0])
	passport := Passport{}

	json.Unmarshal(passportAsBytes, &passport)
	passport.Name = args[1]

	passportAsBytes, _ = json.Marshal(passport)
	APIstub.PutState(args[0], passportAsBytes)

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
