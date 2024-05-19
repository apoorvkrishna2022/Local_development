package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	// This channel receives a value after 2 seconds

	fmt.Println("After 2 seconds")

	fmt.Println("End")

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Execution is finished")
	default:
		fmt.Println("not executed time.After")
	}
}
