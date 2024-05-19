package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(c chan int, somevalue int, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println(somevalue, " has been inserted into the channel")

	c <- somevalue
	//fmt.Println("inside function ", somevalue, " ", time.Now().UnixNano())

	//if somevalue == 9 {
	//	fmt.Println("closing the channel")
	//	close(c)
	//}
}

func main() {

	var wg sync.WaitGroup
	fooValue := make(chan int, 11)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooValue, i, &wg)
	}
	wg.Wait()
	close(fooValue)
	//time.Sleep(5 * time.Second)
	//close(fooValue)
	sum := 0
	for item := range fooValue {
		fmt.Println("adding values", " ", item, " ", time.Now().UnixNano())
		sum += item
	}

	fmt.Println("final sum", sum)
}
