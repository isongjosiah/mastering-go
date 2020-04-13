package main

import (
	"fmt"
)

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	i := -10
	j := 25

	pI := &i
	pJ := &j

	fmt.Println("PI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	*pI = 123456
	fmt.Println("i:", i)
	*pI--
	fmt.Println("i:", i)

	fmt.Println("J:", j)
	getPointer(pJ)
	fmt.Println("j:", j)
	k := returnPointer(12)
	fmt.Println(*k)
	fmt.Println(k)
}
