package main

import (
	"fmt"
	"time"
)

func producer1() int {
	fmt.Printf("%s - enter producer1 function\n", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Printf("%s - exit producer1 function\n", time.Now())

	return 1
}

func producer2() int {
	fmt.Printf("%s - enter producer2 function\n", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Printf("%s - exit producer2 function\n", time.Now())

	return 2
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	// consumer for channel1
	go func() {
		for {
			v1 := <-channel1
			fmt.Printf("%s - consumer1 receive value %d\n", time.Now(), v1)
		}
	}()

	// consumer for channel2
	go func() {
		for {
			v2 := <-channel2
			fmt.Printf("%s - consumer2 receive value %d\n", time.Now(), v2)
		}
	}()

	for {
		select {
		// producer for channel1
		case channel1 <- producer1():
		// producer for channel2
		case channel2 <- producer2():
		}
		// slow down the speed of producer
		time.Sleep(time.Second)
	}
}
