package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NUMBER_OF_PIZZAS = 10

var pizzasMade, pizzasFailed, total int

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch

	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NUMBER_OF_PIZZAS {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order # %d! \n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making Pizza #%d. It will take %d seconds... \n", pizzaNumber, delay)

		// delay for a bit

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** we ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready !", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {

	//keep track of which pizza we are making

	var i = 0

	// want to run forever until we receive a quit notification
	//try to make pizzas

	for {

		currentPizza := makePizza(i)
		// try to make a pizza
		// decision
	}

}

func main() {

	// seed the random number generator

	rand.Seed(time.Now().UnixNano())

	// print out a message

	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background

	go pizzeria(pizzaJob)

	// create and run consumer

	// print out the ending message
}
