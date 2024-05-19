package main

import (
	"fmt"
	"time"
)

func inputIntoChannel(myChannel chan int) {
	time.Sleep(2 * time.Second)
	myChannel <- 4
	return

}
func main() {
	myChannel := make(chan int)

	go inputIntoChannel(myChannel)

	value := <-myChannel

	fmt.Println("Value form the channel ", value)

}
