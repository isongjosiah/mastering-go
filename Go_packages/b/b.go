package b

import (
	"fmt"

	"a"
)

func init() {
	fmt.Println("init() b")
}

// FromB ...
func FromB() {
	fmt.Println("formB()")
	a.FromA()
}
