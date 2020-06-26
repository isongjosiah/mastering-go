package main

import (
	"fmt"
	"time"
)

func f1(c chan int, x int) {
	fmt.Println(x)
	x = <-c
	time.Sleep(1 * time.Second)
	close(c)
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
	//  cannot use x = <- c here to recieve from a send only cahnnel
	c <- x
	time.Sleep(1 * time.Second)
	close(c)
}

func main() {
	c := make(chan int)
	go f1(c, 10)
	time.Sleep(2 * time.Second)
	k := <-c
	time.Sleep(1 * time.Second)
	fmt.Println(k)
	time.Sleep(1 * time.Second)
}
