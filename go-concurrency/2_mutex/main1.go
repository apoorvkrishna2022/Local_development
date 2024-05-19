package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {

	var wg sync.WaitGroup
	var mutex sync.Mutex

	// variable for bank balance

	var bankBalance int

	//print out starting values

	fmt.Printf("Initial Account Balance: %d", bankBalance)
	fmt.Println()

	//define weekly revenue

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Giftes", Amount: 10},
		{Source: "Part Time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	// loop through 52 weeks and print out how much is made; keep a running total

	for i, income := range incomes {

		wg.Add(1)

		go func(i int, income Income, wg *sync.WaitGroup, m *sync.Mutex) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				m.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				m.Unlock()

				fmt.Printf("on Week %d, you earned %d form %s\n", week, income.Amount, income.Source)
			}

		}(i, income, &wg, &mutex)

	}
	wg.Wait()

	// print out final balance

	fmt.Printf("Final bank balance: %d", bankBalance)
	fmt.Println()

}
