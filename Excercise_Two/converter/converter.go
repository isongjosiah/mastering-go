package main

import "fmt"

func main() {
	// a more robust application could recive numbers from the cli and then create an array with that number and convert it to a slice. would work on this later.
	number := [5]int{1, 2, 3, 4, 5}

	slices := number[:]
	iMap := make(map[int]int)

	for key, value := range slices {
		iMap[key] = value
	}

	fmt.Println("Array succesfully converted to a map")
	fmt.Printf("Previous array: %d\n", number)
	fmt.Printf("Converted map: %d\n", iMap)
}
