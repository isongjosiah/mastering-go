package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("k", true, "k flag")
	minusD := flag.Int("O", 1, "O")
	flag.Parse()

	valueK := *minusK
	valueO := *minusD
	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)
}
