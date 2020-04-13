package main

import (
	"fmt"
)

func main() {
	// typical for loop
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, "")
	}

	fmt.Println()
	// typical while loop
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	// typical do...while loop
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// using range
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index: ", i, "value: ", value)
	}
}
