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
	Password     string  `json:"password"`
	Image        string  `json:"image"`
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
	} else if function == "queryPassportsByPassNb" { //find marbles for owner X using rich query
		return s.queryPassportsByPassNb(APIstub, args)
	} else if function == "validNumPwd" { //find marbles for owner X using rich query
		return s.validNumPwd(APIstub, args)
	} else if function == "changePassportOwner" {
		return s.changePassportOwner(APIstub, args)
	} else if function == "changePassport" {
		return s.changePassport(APIstub, args)
	} else if function == "querykeybyPassNb" {
		return s.querykeybyPassNb(APIstub, args)
	} else if function == "searchPassportByCountry" {
		return s.searchPassportByCountry(APIstub, args)
	} else if function == "changePassportValidity" {
		return s.changePassportValidity(APIstub, args)
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
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML52147", Name: "Jean", Surname: "Dupont", DateOfBirth: "16/09/1985", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "25/11/2013", PassOrigin: "France", Validity: "Valide", Password: "4acc06cf6378d890167841c340fdfc69e8e9bde8", Image: "testimage"},
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML22389", Name: "Brad", Surname: "Dupont", DateOfBirth: "10/03/1975", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "5/07/2017", PassOrigin: "France", Validity: "Valide", Password: "4acc06cf6378d890167841c340fdfc69e8e9bde8", Image: "testimage"},
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML66146", Name: "Jin Soo", Surname: "Dupont", DateOfBirth: "1/05/2000", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "2/01/2015", PassOrigin: "France", Validity: "Valide", Password: "4acc06cf6378d890167841c340fdfc69e8e9bde8", Image: "testimage"},
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

	if len(args) != 19 {
		return shim.Error("Incorrect number of arguments. Expecting 18")
	}

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	var buffer bytes.Buffer
	var i int

	i = 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		buffer.WriteString(queryResponse.Key)

		i = i + 1
	}

	taille, _ := strconv.ParseFloat(args[9], 64)

	queryString := fmt.Sprintf("{\"selector\":{\"passNb\":\"%s\"}}", args[2])

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	if queryResults == nil {
		var passport = Passport{Type: args[0], CountryCode: args[1], PassNb: args[2], Name: args[3], Surname: args[4], DateOfBirth: args[5], Nationality: args[6], Sex: args[7], PlaceOfBirth: args[8], Height: taille, Autority: args[10], Residence: args[11], EyesColor: args[12], DateOfExpiry: args[13], DateOfIssue: args[14], PassOrigin: args[15], Validity: args[16], Password: args[17], Image: args[18]}
		passportAsBytes, _ := json.Marshal(passport)
		APIstub.PutState(strconv.Itoa(i), passportAsBytes)
		return shim.Success(nil)
	} else {
		return shim.Error(err.Error())
	}

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

	query, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	buffer.Write(query.Bytes())
	buffer.WriteString("]")

	fmt.Printf("- queryAllPassports:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryPassportsByPassNb(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	PassNb := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"passNb\":\"%s\"}}", PassNb)

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
func (s *SmartContract) searchPassportByCountry(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	CountryCode := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"countryCode\":\"%s\"}}", CountryCode)

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.Write(queryResults)
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) validNumPwd(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	var buffer bytes.Buffer

	PassNb := args[0]
	Pwd := args[1]
	queryString := fmt.Sprintf("{\"selector\":{\"passNb\":\"%s\",\"password\":\"%s\"}}", PassNb, Pwd)

	queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	if queryResults != nil {
		buffer.WriteString("true")
		return shim.Success(buffer.Bytes())
	}
	buffer.WriteString("false")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changePassportValidity(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	key, err := getKeybyPAssnum(APIstub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	passportAsBytes, _ := APIstub.GetState(key)
	passport := Passport{}

	json.Unmarshal(passportAsBytes, &passport)
	if passport.Validity == "Invalide" {
		passport.Validity = "Valide"
	} else if passport.Validity == "Valide" {
		passport.Validity = "Invalide"
	}

	passportAsBytes, _ = json.Marshal(passport)
	APIstub.PutState(key, passportAsBytes)

	return shim.Success(nil)
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

func (s *SmartContract) querykeybyPassNb(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 19")
	}

	key, err := getKeybyPAssnum(APIstub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	passportAsBytes, _ := APIstub.GetState(key)
	passport := Passport{}

	json.Unmarshal(passportAsBytes, &passport)
	var buffer bytes.Buffer
	buffer.WriteString(key + ":" + passport.PassNb)
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changePassport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 19 {
		return shim.Error("Incorrect number of arguments. Expecting 19")
	}

	key, err := getKeybyPAssnum(APIstub, args[2])
	if err != nil {
		return shim.Error(err.Error())
	}
	passportAsBytes, _ := APIstub.GetState(key)
	passport := Passport{}

	json.Unmarshal(passportAsBytes, &passport)
	taille, _ := strconv.ParseFloat(args[9], 64)
	passport.Type = args[0]
	passport.CountryCode = args[1]
	passport.Name = args[3]
	passport.Surname = args[4]
	passport.DateOfBirth = args[5]
	passport.Nationality = args[6]
	passport.Sex = args[7]
	passport.PlaceOfBirth = args[8]
	passport.Height = taille
	passport.Autority = args[10]
	passport.Residence = args[11]
	passport.EyesColor = args[12]
	passport.DateOfExpiry = args[13]
	passport.DateOfIssue = args[14]
	passport.PassOrigin = args[15]
	passport.Validity = args[16]
	passport.Password = args[17]
	passport.Image = args[18]
	passportAsBytes, _ = json.Marshal(passport)
	APIstub.PutState(key, passportAsBytes)
	return shim.Success(nil)
}

func getKeybyPAssnum(APIstub shim.ChaincodeStubInterface, PassNb string) (string, error) {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return "", err
	}
	passport := Passport{}

	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return "", err
		}
		passport = Passport{}
		json.Unmarshal(queryResponse.Value, &passport)
		if passport.PassNb == PassNb {
			return queryResponse.Key, nil
		}

	}
	return "", nil
}

func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
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

	return &buffer, nil
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
