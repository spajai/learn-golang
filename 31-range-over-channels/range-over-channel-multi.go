package main

import "fmt"

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send]:
	// main.main()
	// 	/home/sushrut/go-lang-learn/31-range-over-channels/range-over-channel-multi.go:8 +0x65
	// exit status 2
	// above error when we send max data than buffer
	queue := make(chan int, 50)

	for i := 0; i < 50; i++ {
		queue <- i
	}

	close(queue)
	for elm := range queue {
		fmt.Println(elm)
	}

}
