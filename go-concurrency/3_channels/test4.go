package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generateRandomNumber() int {
	maxRange := 1000
	return rand.Intn(maxRange)
}

func transform(val int) string {
	temp := fmt.Sprintf("The value is %d", val)
	return temp

}

func main() {

	length := 10

	myChannel := make(chan int, length+1)
	for i := 0; i < length; i++ {
		myChannel <- generateRandomNumber()

	}
	close(myChannel)

	myResultChannel := make(chan string, length+1)

	maxNumberOfGoroutines := 8

	var wg sync.WaitGroup

	for maxNumberOfGoroutines > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				data, ok := <-myChannel
				if !ok {
					return
				}
				myResultChannel <- transform(data)
			}
		}()
		maxNumberOfGoroutines--
	}

	wg.Wait()
	close(myResultChannel)

	for ans := range myResultChannel {
		fmt.Println(ans)
	}

}
