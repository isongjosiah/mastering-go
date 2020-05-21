package main

import (
	"fmt"
	"mypackage"
)

func main() {
	fmt.Println("Using mypackage!")
	mypackage.A()
	mypackage.B()
	fmt.Println(mypackage.MyConstant)
}
