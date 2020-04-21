package main

import (
	"encoding/xml"
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

// Decodes XML from a file to a key, returns an error if it encouters one
func loadFromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
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
	err := loadFromXML(filename, &myRecord)

	if err != nil {
		fmt.Println("XML:", myRecord)
	} else {
		fmt.Println(err)
	}
}
