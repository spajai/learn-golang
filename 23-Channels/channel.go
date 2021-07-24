package main

import "fmt"

func main() {
	messages := make(chan string)

	//annonymous sub to send the string to channel
	go func() {
		messages <- fmt.Sprintf("Ping from host")
	}()

	//listen from channel
	msg := <-messages
	fmt.Println(msg)
}

// go func() {
// 	for i := 0; i < 20; i++ {
// 		messages <- fmt.Sprintf("Ping from host no : %d", i)
// 	}
// }()

// for i := 0; i < 20; i++ {
// 	msg := <-messages
// 	fmt.Println(msg)
// }
