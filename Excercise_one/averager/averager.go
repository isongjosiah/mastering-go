package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please enter one or more arguments")
		os.Exit(1)
	}

	err := errors.New("an error")
	var m float64
	k := 1
	arguments := os.Args

	// minus 1 because the name of the program is also an element of os.Args
	// n := float64(len(arguments) - 1) [This would not work, because it is possible to give a non number in combination with a number and we will then count the non-number as a number]

	var n float64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments was a float !")
			return
		}

		m, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	//  we make n one here to account for the first argument
	n = 1
	var sum float64
	sum = m

	for i := 2; i < len(os.Args); i++ {
		m, err = strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += m
			n++
		}
	}

	fmt.Println("Sum :", sum)
	fmt.Println("n :", n)
	average := sum / n

	fmt.Println("Average :", average)
}
