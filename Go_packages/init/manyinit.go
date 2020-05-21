package main

import (
	"a"
	"fmt"

	"b"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
}
