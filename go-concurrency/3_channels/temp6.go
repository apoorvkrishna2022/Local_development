package main

import "fmt"

func main() {

	ch := make(chan int)

	var val int

	go func() {
		val = <-ch

	}()

	ch <- 10

	fmt.Println(val)

	return
}
