package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleIntance *single

func getInstance() *single {
	if singleIntance == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Creating single instance now.")
		singleIntance = &single{}
	} else {
		fmt.Println("signle is created")
	}

	return singleIntance
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
