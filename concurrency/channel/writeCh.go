package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)

	// reading from the channel c
	// k := <-c
	// noticed that after reading from the channel it reverts to the
	// nil value of it's type 0 for int.
	fmt.Println("Read:", <-c)
	// fmt.Println("k:", k)
	time.Sleep(1 * time.Second)

	// Determining if a channel is open or closed.
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed !")
	}
}
