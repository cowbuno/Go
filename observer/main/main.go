package main

import (
	"SDP/observer"
	"fmt"
)

func main() {
	firstObs := observer.Customer{name: "bob", email: "bob@"}
	fmt.Println(firstObs)
}
