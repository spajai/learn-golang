package main

import (
	"fmt"
)

func main() {

	// if no case match default will be executed

	messages := make(chan string)
	signals := make(chan bool)
	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	select {
	case msg := <-messages:
		fmt.Println("Received", msg)
	default:
		fmt.Println("no messages received")
	}

	msg := "hi"
	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.
	select {
	case messages <- msg:
		fmt.Println("sent ", messages)
	default:
		fmt.Println("no messages sent")
	}
	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.

	select {
	case msg := <-messages:
		fmt.Println("Received messaged", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
