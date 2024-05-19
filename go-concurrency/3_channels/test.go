package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	go greet1(msg)
	go greet2(msg)

	for {
		greeting1 := <-msg
		fmt.Println("Greeting received")
		fmt.Println(greeting1)

		greeting2, ok := <-msg
		if ok {
			fmt.Println("Channel is open!")

			fmt.Println("Greeting received")
			fmt.Println(greeting2)
		} else {
			fmt.Println("Channel is closed!")
		}

	}

	// Close the channel outside of the goroutines
	close(msg)
}

func greet1(ch chan string) {
	fmt.Println("Greeter  1 waiting to send greeting!")
	time.Sleep(2 * time.Second)
	ch <- "Hello Rwitesh"
	fmt.Println("Greeter 1 completed")
}

func greet2(ch chan string) {
	fmt.Println("Greeter 2 waiting to send greeting!")
	time.Sleep(5 * time.Second)
	ch <- "Hello Rwitesh 2"
	fmt.Println("Greeter 2 completed")
}
