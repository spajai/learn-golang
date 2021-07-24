package main

import "fmt"

func main() {

	//  length is 3
	messages := make(chan string, 3)

	// send messages 3 time
	messages <- "Buffered"
	messages <- "Channel"
	messages <- "test"

	// print 3 times
	// size and sending message mismatch can cause the deadlock
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
