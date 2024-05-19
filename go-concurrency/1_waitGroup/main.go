package main

import (
	"fmt"
	"sync"
)

func printStirng(s string, wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}

	fmt.Println(s)
}

//func main() {
//
//	go printStirng("first string")
//
//	time.Sleep(1 * time.Second)
//
//	printStirng("second string")
//}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	//wg.Add(9)

	printStirng("first string", nil)
	for i, x := range words {
		go printStirng(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	printStirng("second string", nil)

	wg.Wait()

	printStirng("third string", nil)

}
