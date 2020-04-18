package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

// Global variable to handle all the IPFS API client calls
var sh *shell.Shell

func main() {
	// key-value pair
	keyValueMap := make(map[string]interface{})

	// localnode: infura (change for localhost)
	sh = shell.NewShell("https://ipfs.infura.io:5001")

	// print
	fmt.Println("This client generates a dynamic key-value entry and stores it in IPFS!")

	// scan
	inputKey, inputValue := ScanInput()

	// set
	keyValueMap[inputKey] = inputValue

	// Converting into JSON object
	entryJSON, err := json.Marshal(keyValueMap)
	if err != nil {
		fmt.Println(err)
	}

	// Display the marshaled JSON object before sending it to IPFS
	jsonStr := string(entryJSON)
	fmt.Println("The JSON object of your key-value entry is:")
	fmt.Println(jsonStr)

	// Dag PUT operation which will return the CID for futher access or pinning etc.
	cid, err := sh.DagPut(entryJSON, "json", "cbor")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println("------\nOUTPUT\n------")
	fmt.Printf("WRITE: Successfully added %sHere's the IPLD Explorer link: https://explore.ipld.io/#/explore/%s \n", string(cid+"\n"), string(cid+"\n"))

	// Fetch the details by reading the DAG for key "inputKey"
	fmt.Printf("READ: Value for key \"%s\" is: ", inputKey)
	res, err := GetDag(cid, inputKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}

// GetDag handles READ operations of a DAG entry by CID, returning the corresponding value
func GetDag(ref, key string) (out interface{}, err error) {
	err = sh.DagGet(ref+"/"+key, &out)
	return
}

func ScanInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter value for the key field: ")
	scanner.Scan()
	inputKey := scanner.Text()

	fmt.Println("Enter value for the value field: ")
	scanner.Scan()
	inputValue := scanner.Text()

	return inputKey, inputValue
}
