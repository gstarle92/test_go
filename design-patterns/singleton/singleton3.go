package main

import "sync"

var once sync.Once

type singleton2 map[string]string

var (
	instance2 singleton2
)

func NewInctance2() singleton2 {

	once.Do(func() {
		instance2 = make(singleton2)
	})

	return instance2
}

func main() {
	NewInctance2()
}
