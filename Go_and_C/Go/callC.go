package main

/*
#cgo CFLAGS: -I${SRCDIR}/C
#cgo LDFLAGS: ${SRCDIR}/callC.a
#include <stdlib.h>
#include <callC.h>
*/
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println("Going to call a C function !")
	C.cHello()
	fmt.Println("Going to call another c function!")
	myMessage := C.Cstring
}
