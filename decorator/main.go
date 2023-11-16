package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 3
}

type CheeseTopping struct {
	pizza IPizza
}

func (t *CheeseTopping) getPrice() int {
	return t.pizza.getPrice() + 2
}

func main() {
	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	pizzaWithDoubleCheese := &CheeseTopping{
		pizza: pizzaWithCheese,
	}
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithDoubleCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and doublecheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
