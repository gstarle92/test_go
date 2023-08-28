package main

import (
	"fmt"
	"sync"
)

var X = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {

	wg.Add(2)
	go increment1(&wg, &mutex)
	go increment2(&wg, &mutex)

	wg.Wait()
	fmt.Println("Final Value of X is", X)
}

func increment1(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	X = X + 1
	mutex.Unlock()
	fmt.Println("increment1 has incremented value of x", X)
}

func increment2(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	X = X + 1
	mutex.Unlock()
	fmt.Println("increment2 has incremented value of x", X)
}
