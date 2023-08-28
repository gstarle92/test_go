package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

type DbConnection struct {
	conn string
}

var DBInstance *DbConnection

func NewDBConnection(mutex *sync.Mutex, wg *sync.WaitGroup) *DbConnection {
	defer wg.Done()
	mutex.Lock()
	if DBInstance == nil {
		DBInstance = &DbConnection{
			conn: "this is mysql Connection ",
		}
		fmt.Println("instance Initialized")
	} else {
		fmt.Println("instance already Initialized")
	}
	mutex.Unlock()
	return DBInstance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go NewDBConnection(&mutex, &wg)
	}
	wg.Wait()

	//fmt.Println(DBInstance.conn)
}

// func connectionFun(mutex sync.Mutex, wg sync.WaitGroup) {

// 	DBInstance = NewDBConnection()
// 	fmt.Println(DBInstance.conn)
// 	wg.Done()
// }
