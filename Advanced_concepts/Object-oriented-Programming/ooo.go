package main

import "fmt"

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

// A ...A
func (A a) A() {
	fmt.Println("Function A() for A")
}

// B ...
func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var i c
	i.A.A()
	i.B.A()
}
