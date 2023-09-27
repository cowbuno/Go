package main

import (
	"SDP/strategy"
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 1, 7, 9}

	context1 := strategy.NewContext(&strategy.BubbleSort{})
	fmt.Println(arr)
	arr = context1.Sort(arr)
	fmt.Println(arr)

	arr = []int{5, 4, 3, 1, 7, 9}
	context2 := strategy.NewContext(&strategy.QuickSort{})
	fmt.Println(arr)
	arr = context2.Sort(arr)
	fmt.Println(arr)

}
