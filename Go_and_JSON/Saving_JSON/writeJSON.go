package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Record ...
type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

// Telephone ...
type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel:     []Telephone{{Mobile: true, Number: "1234-567"}, {Mobile: true, Number: "1234-abcd"}, {Mobile: false, Number: "abcd-567"}},
	}
	saveToJSON(os.Stdout, myRecord)
}
