package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// Employee struct to store pension account details on the ledger.
type Employee struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Age                 int            `json:"age"`
	Employer            string         `json:"employer"`
	ContributionHistory []Contribution `json:"contributionHistory"`
	Balance             float64        `json:"balance"`
}

// Contribution struct to record individual contributions.
type Contribution struct {
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

// Retirement age constant
const retirementAge = 60

// Init is called when the smart contract is instantiated
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("pension_cc")

// Invoke is called per transaction on the chaincode.
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %s", function)
	logger.Infof("Args length is : %d", len(args))

	switch function {
	case "GetEmployee":
		return s.GetEmployee(APIstub, args)
	case "InitLedger":
		return s.InitLedger(APIstub)
	case "CreateEmployee":
		return s.CreateEmployee(APIstub, args)
	case "GetAllEmployees":
		return s.GetAllEmployees(APIstub)
	case "ContributeToPension":
		return s.ContributeToPension(APIstub, args)
	case "WithdrawPension":
		return s.WithdrawPension(APIstub, args)
	case "GetEmployeeHistory":
		return s.GetEmployeeHistory(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}
}

// GetEmployee queries the ledger for an employee by ID.
func (s *SmartContract) GetEmployee(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 (EmployeeID)")
	}

	employeeAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if employeeAsBytes == nil {
		return shim.Error("Employee not found: " + args[0])
	}

	return shim.Success(employeeAsBytes)
}

// InitLedger adds a base set of employees to the ledger.
func (s *SmartContract) InitLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	employees := []Employee{
		{ID: "EMP001", Name: "Alice Johnson", Age: 34, Employer: "TechCorp", Balance: 15000.00},
		{ID: "EMP002", Name: "Bob Williams", Age: 58, Employer: "FinanceInc", Balance: 75000.00},
		{ID: "EMP003", Name: "Charlie Brown", Age: 61, Employer: "HealthCo", Balance: 120000.00},
	}

	for _, employee := range employees {
		employeeJSON, err := json.Marshal(employee)
		if err != nil {
			return shim.Error("Failed to marshal employee: " + err.Error())
		}
		err = APIstub.PutState(employee.ID, employeeJSON)
		if err != nil {
			return shim.Error("Failed to put employee to world state: " + err.Error())
		}
	}

	return shim.Success(nil)
}

// CreateEmployee creates a new employee record on the ledger.
func (s *SmartContract) CreateEmployee(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5 (EmployeeID, Name, Age, Employer, ContributionAmount)")
	}

	// Check if employee already exists
	existing, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get employee: " + err.Error())
	}
	if existing != nil {
		return shim.Error("Employee already exists: " + args[0])
	}

	age, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid age. Must be an integer.")
	}

	contributionAmount, err := strconv.ParseFloat(args[4], 64)
	if err != nil {
		return shim.Error("Invalid contribution amount. Must be a number.")
	}

	var employee = Employee{
		ID:                  args[0],
		Name:                args[1],
		Age:                 age,
		Employer:            args[3],
		ContributionHistory: []Contribution{},
		Balance:             0.0,
	}

	if contributionAmount > 0 {
		contribution := Contribution{
			Amount:    contributionAmount,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}
		employee.ContributionHistory = append(employee.ContributionHistory, contribution)
		employee.Balance = contributionAmount
	}

	employeeAsBytes, _ := json.Marshal(employee)
	err = APIstub.PutState(args[0], employeeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(employeeAsBytes)
}

// ContributeToPension adds a contribution to an employee's pension fund.
func (s *SmartContract) ContributeToPension(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 (EmployeeID, Amount)")
	}

	employeeID := args[0]
	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return shim.Error("Invalid contribution amount. Must be a number.")
	}
	if amount <= 0 {
		return shim.Error("Contribution amount must be positive.")
	}

	employeeAsBytes, err := APIstub.GetState(employeeID)
	if err != nil {
		return shim.Error(err.Error())
	}
	if employeeAsBytes == nil {
		return shim.Error("Employee not found: " + employeeID)
	}

	var employee Employee
	json.Unmarshal(employeeAsBytes, &employee)

	contribution := Contribution{
		Amount:    amount,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	employee.ContributionHistory = append(employee.ContributionHistory, contribution)
	employee.Balance += amount

	updatedEmployeeAsBytes, _ := json.Marshal(employee)
	err = APIstub.PutState(employeeID, updatedEmployeeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(updatedEmployeeAsBytes)
}

// WithdrawPension allows an employee to withdraw their funds if they are of retirement age.
func (s *SmartContract) WithdrawPension(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 (EmployeeID)")
	}

	employeeID := args[0]
	employeeAsBytes, err := APIstub.GetState(employeeID)
	if err != nil {
		return shim.Error(err.Error())
	}
	if employeeAsBytes == nil {
		return shim.Error("Employee not found: " + employeeID)
	}

	var employee Employee
	json.Unmarshal(employeeAsBytes, &employee)

	if employee.Age < retirementAge {
		return shim.Error(fmt.Sprintf("Employee %s is not eligible for withdrawal. Retirement age is %d, current age is %d", employeeID, retirementAge, employee.Age))
	}

	if employee.Balance <= 0 {
		return shim.Error("Employee has no balance to withdraw.")
	}

	// For simplicity, we set the balance to 0. A real system might move it to a "paid" state.
	employee.Balance = 0.0

	updatedEmployeeAsBytes, _ := json.Marshal(employee)
	err = APIstub.PutState(employeeID, updatedEmployeeAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(updatedEmployeeAsBytes)
}

// GetAllEmployees returns all employee records from the ledger.
func (s *SmartContract) GetAllEmployees(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Using a key prefix to fetch all employees
	startKey := "EMP0"
	endKey := "EMP9999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

// GetEmployeeHistory returns the modification history for a given employee.
func (s *SmartContract) GetEmployeeHistory(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 (EmployeeID)")
	}

	employeeID := args[0]
	logger.Infof("Getting history for EmployeeID: %s", employeeID)

	resultsIterator, err := stub.GetHistoryForKey(employeeID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
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
		buffer.WriteString(strconv.FormatBool(response.IsDelete))

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode.
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
