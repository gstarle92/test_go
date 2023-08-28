package main

import "sync"

var mutex1 sync.Mutex

type singleton1 map[string]string

var (
	instance1 singleton1
)

func NewSingleton1() singleton1 {
	mutex1.Lock()
	defer mutex1.Unlock()
	if instance1 == nil {
		instance1 = make(singleton1)
	}
	return instance1
}
func main() {
	NewSingleton1()
}
