package main

import (
	"fmt"
	"time"
)

func main() {
	// create two channel
	c1 := make(chan string)
	c2 := make(chan string)

	// go routine

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// We’ll use select to await both of these values
		// simultaneously, printing each one as it arrives.
		//select will wait

		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

}

// wont work select expect the case statement
// select {
// 	if <-c1 {
// 		fmt.Println("Received", j)
// 	}else if <-c2 {
// 		fmt.Println("Received", k)
// 	}
// }
