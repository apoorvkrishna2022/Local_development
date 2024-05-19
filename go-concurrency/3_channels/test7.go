package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int, value int, delay time.Duration) {
	time.Sleep(delay)
	ch <- value
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//go sendData(ch1, 10, 2*time.Second)
	//go sendData(ch2, 20, 3*time.Second)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("print time from ch1 ", time.Now())
		fmt.Println("Received from ch1:", <-ch1)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received from ch2:", <-ch2, " ", time.Now().UnixNano())
	}()

	select {

	case ch1 <- 30:
		fmt.Println("print time to ch1 ")
		fmt.Println("Sent data 30 to ch1")
	case ch2 <- 40:
		fmt.Println("print time to ch2 ", time.Now().UnixNano())
		fmt.Println("Sent data 40 to ch2 ", time.Now().UnixNano())
	}

	// Wait to receive data from channels (in a real-world scenario, you may need to do this)
	fmt.Println("existing from main")
}
