package mypackage

import "fmt"

// A ...
func A() {
	fmt.Println("This is function A!")
}

// B ...
func B() {
	fmt.Println("privateConstant:", privateConstant)
}

// MyConstant ...
const MyConstant = 123
const privateConstant = 21
