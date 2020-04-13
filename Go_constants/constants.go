package main

import (
	"fmt"
)

// Digit ...
type Digit int

// Power2 ...
type Power2 int

const pi = 3.1415926

const (
	c1 = "C1C1C1"
	c2 = "C2C2C2"
	c3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(pi)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
