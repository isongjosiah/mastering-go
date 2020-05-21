package main

import "fmt"

// *T is a pointer to a space in memory that contains a value of type T
func getPtr(v *float64) float64 {
	fmt.Print(v)
	return *v * *v
}

func returnPtr(x int) *int {
	y := x * x
	return &y
}

func main() {
	x := 12.2
	fmt.Println(&x)
	fmt.Println(getPtr(&x))
	x = 12.0
	fmt.Println(getPtr(&x))

	sq := returnPtr(10)
	fmt.Println("sq value:", *sq)
	fmt.Println("sq memory address:", sq)
}
