package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		ch2 <- 2
	}()

	//<-ch1
	//<-ch2

	//time.Sleep(4 * time.Second)
	for i := 0; i < 2; i++ {
		fmt.Println("i", i)
		//Loop:
		//for {
		select {
		case <-ch1:
			fmt.Println("Received from ch1")
			//break Loop // break out of the labeled loop
		case <-ch2:
			fmt.Println("Received from ch2")
			//break Loop // break out of the labeled loop
			//default:
			//	//fmt.Println("Default case executed")
		}
		//}
	}

	wg.Wait()

	fmt.Println("Leaving main thread")
}
