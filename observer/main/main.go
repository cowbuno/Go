package main

import (
	"SDP/observer"
)

func main() {
	obs1 := observer.NewCustomer("Bob", "bob@gmail.com")
	obs2 := observer.NewCustomer("Alice", "alice@gmail.com")
	obs3 := observer.NewCustomer("obs3", "obs3@gmail.com")

	iphone_15 := observer.NewItem("Iphone 15")

	iphone_15.addObserver(obs1)
	iphone_15.addObserver(obs2)
	iphone_15.addObserver(obs3)

	iphone_15.printAllObserver()

}
