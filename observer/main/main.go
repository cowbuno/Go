package main

import (
	"SDP/observer"
)

func main() {
	iphone_15 := observer.NewItem("Iphone 15")

	c1 := observer.NewCustomer("Bob", "bob@gmail.com")
	c2 := observer.NewCustomer("Alice", "alice@gmail.com")
	c3 := observer.NewCustomer("obs3", "obs3@gmail.com")

	iphone_15.AddObserver(c1)
	iphone_15.AddObserver(c2)
	iphone_15.AddObserver(c3)

	iphone_15.PrintAllObserver()

}
