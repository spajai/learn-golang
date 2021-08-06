package main

import "fmt"

func main() {
	// In this example we'll use a `jobs` channel to
	// communicate work to be done from the `main()` goroutine
	// to a worker goroutine. When we have no more jobs for
	// the worker we'll `close` the `jobs` channel.
	jobs := make(chan int, 5)
	done := make(chan bool)
	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.
	go func() {
		for {
			j, more := <-jobs
			// 3 true
			// fmt.Println(j, more)
			if more {
				fmt.Println("Received Job", j)
			} else {
				fmt.Println("Received all Jobs")
				done <- true
				return
			}

		}

	}()

	// if sent on closed channel
	// panic: send on closed channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent Job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	<-done

}
