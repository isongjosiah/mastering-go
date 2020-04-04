package main

import (
	"fmt"
	"os"
	"strconv"
)

// strconv helps convert a string into an arithmetical data type

func main() {
	// The first element in the os.Args slice is the name of the executable program
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
