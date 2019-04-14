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
 * The Init method is called when the Smart Contract  is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
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
	testimage := "data:image/jpeg;base64,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"

	passports := []Passport{
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML52147", Name: "Jean", Surname: "Dupont", DateOfBirth: "16/09/1985", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "25/11/2013", PassOrigin: "France", Validity: "Valide", Password: "4acc06cf6378d890167841c340fdfc69e8e9bde8", Image: testimage},
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML22389", Name: "Brad", Surname: "Dupont", DateOfBirth: "10/03/1975", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "5/07/2017", PassOrigin: "France", Validity: "Valide", Password: "4acc06cf6378d890167841c340fdfc69e8e9bde8", Image: testimage},
		Passport{Type: "P", CountryCode: "FR", PassNb: "14ML66146", Name: "Jin Soo", Surname: "Dupont", DateOfBirth: "1/05/2000", Nationality: "France", Sex: "M", PlaceOfBirth: "Toulouse", Height: taille, Autority: "Préfecture de ", Residence: "Avenue des Facultés, 33400 Talence", EyesColor: "Marron", DateOfExpiry: "16/02/2023", DateOfIssue: "2/01/2015", PassOrigin: "France", Validity: "Valide", Password: "4acc06cf6378d890167841c340fdfc69e8e9bde8", Image: testimage},
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

	if len(args) != 18 {
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
	passport.Image = args[17]
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
