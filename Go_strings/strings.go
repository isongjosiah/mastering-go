package main

import (
	"fmt"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x", sLiteral[i])
	}
	fmt.Println()
	// prints a double-quoted string that is safely escaped with Go syntax
	fmt.Printf("q: %q\n", sLiteral)
	// guarantees ASCII-only output
	fmt.Printf("+q: %+q\n", sLiteral)
	// puts spaces between the printed bytes
	fmt.Printf(" x: % x\n", sLiteral)
	// in order to print a string literal as a string you willl need to call fmt.Printf() with %s
	fmt.Printf("s: As a string: %s\n", sLiteral)

	s2 := "€£³"

	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	fmt.Printf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: %x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x", s3[i])
	}

	fmt.Println()

}
