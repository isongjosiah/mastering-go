package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var myInt int64
	arguments := os.Args
	err := errors.New("an error")
	k := 1

	if len(arguments) == 1 {
		fmt.Println("Please give me one argument!")
		os.Exit(1)
	}

	//  error checking to make sure there is at least an integer
	for err != nil {
		if k > len(arguments) {
			fmt.Println("Please enter a valid integer")
			return
		}

		if k < len(arguments) {
			myInt, err = strconv.ParseInt(arguments[k], 10, 64)
			k++
		}
	}
	// Just to use myInt
	myInt--

	for k := 1; k < len(arguments); k++ {

		_, err = strconv.ParseInt(arguments[k], 10, 64)
		// writes the arguments and stops if it encounters END
		if arguments[k] != "END" {
			if err == nil {
				io.WriteString(os.Stdout, arguments[k])
				io.WriteString(os.Stdout, "\n")
			}
		} else {
			return
		}
	}
}
