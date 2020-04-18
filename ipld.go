package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

type SampleStruct struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Salary string `json:"Salary"`
}

// Global variable to handle all the IPFS API client calls
var sh *shell.Shell

func main() {
	// key-value pair
	DocStoreMap := make(map[string]SampleStruct)

	// localnode: infura (change for localhost)
	sh = shell.NewShell("https://ipfs.infura.io:5001")

	// print
	fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!")

	// scan
	employee := ScanInput()

	// set
	DocStoreMap[employee.ID] = employee

	// Converting into JSON object
	entryJSON, err := json.Marshal(DocStoreMap)
	if err != nil {
		fmt.Println(err)
	}

	// Display the marshaled JSON object before sending it to IPFS
	jsonStr := string(entryJSON)
	fmt.Println("The JSON object of your key-value entry is:")
	fmt.Println(jsonStr)

	// Dag PUT operation which will return the CID for futher access or pinning etc.
	start := time.Now()
	cid, err := sh.DagPut(entryJSON, "json", "cbor")
	elapsed := time.Since(start)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// Outputs
	fmt.Println("------\nOUTPUT\n------")
	fmt.Printf("WRITE: Successfully added %sHere's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s", string(cid+"\n"), string(cid+"\n"))
	fmt.Println("WRITE: IPLD PUT call took ", elapsed)

	// READ and get document
	fmt.Printf("READ: Reading the document details of employee by ID: \"%s\"\n", employee.ID)
	start = time.Now()
	document, err := GetDocument(cid, employee.ID)
	elapsed = time.Since(start)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("READ: Salary of employee ID %s is %s\n", string(employee.ID), string(document.Salary))
	fmt.Println("READ: IPLD GET call took ", elapsed)
}

// GetDocument READ operations of a DAG entry by CID, returning the corresponding value
func GetDocument(ref, key string) (out SampleStruct, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}

func ScanInput() SampleStruct {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the ID of the Employee: ")
	scanner.Scan()
	inputID := scanner.Text()

	fmt.Println("Enter the name of the Employee: ")
	scanner.Scan()
	inputName := scanner.Text()

	fmt.Println("Enter the salary of the Employee: ")
	scanner.Scan()
	inputSalary := scanner.Text()

	employee := SampleStruct{
		ID:     inputID,
		Name:   inputName,
		Salary: inputSalary,
	}

	return employee
}
