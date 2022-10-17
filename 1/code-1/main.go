package main

import (
	"fmt"
	"time"
)

func producer1() int {
	fmt.Printf("%s - enter producer1 function\n", time.Now())

	return 1
}

func main() {
	channel1 := make(chan int)

	// consumer for channel1
	go func() {
		for {
			// sleep 5 seconds to delay the readiness of channel1
			time.Sleep(5 * time.Second)

			v1 := <-channel1
			fmt.Printf("%s - consumer1 receive value %d\n", time.Now(), v1)
		}
	}()

	for {
		select {
		// producer for channel1
		case channel1 <- producer1():
		}
	}
}
