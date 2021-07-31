package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	fmt.Println("Times after 2 sec ", time.After(2*time.Second))
	fmt.Println("Time Seconds", time.Second)

	// GO routine
	// send the messages after 2 second
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	// select the result
	// if the the time reaches 1 second time out
	// above messages will come after 1 sec so timeout
	select {
	case res := <-c1:
		fmt.Println("Received", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	// make another channel
	c2 := make(chan string)
	// wait for 2 second
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println("Received", res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}

}
