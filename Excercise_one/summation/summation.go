package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please give me one or more number.")
		os.Exit(1)
	}

	arguments := os.Args
	k := 1
	var err error = errors.New("an error")
	var m int64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a number")
			return
		}

		m, err = strconv.ParseInt(arguments[k], 10, 64)
		k++
	}
	a := m

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseInt(arguments[i], 10, 64)
		if err == nil {
			a += n
		}
	}

	fmt.Println("Sum: ", a)
}
