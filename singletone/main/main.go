package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s1.data = 5

	s2 := GetInstance()
	fmt.Println(s2.data)
	if s1 == s2 {
		fmt.Println("Both are the same instance.")
	}
}
