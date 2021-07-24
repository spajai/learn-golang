package main

// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,

import "fmt"
import "time"

// `done` channel will be used to notify another
// goroutine that this function's work is done.
// create func woker which takes input of type chan
func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {
	// make channel of type bool
	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	// using thred call worker
	go worker(done)
	// Block until we receive a notification from the worker on the channel.
	<-done
}
