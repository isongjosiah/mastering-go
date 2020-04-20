package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename")
		return
	}

	filename := arguments[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parseData map[string]interface{}
	json.Unmarshal([]byte(fileData), &parseData)

	for key, value := range parseData {
		fmt.Println("key:", key, "value:", value)
	}
}
