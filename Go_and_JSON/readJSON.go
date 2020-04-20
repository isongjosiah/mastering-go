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

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	fmt.Println("Loaded from json")

	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
