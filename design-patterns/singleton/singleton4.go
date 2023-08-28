package main

import (
	"fmt"
	"sync"
)

var once1 sync.Once

type DBConnection struct {
	conn string
}

var DBInstance *DBConnection

func NewDBConnection() *DBConnection {

	// if DBInstance == nil {
	// 	DBInstance = &DBConnection{
	// 		conn: "This is Mysql Connection sting ",
	// 	}
	// }
	once1.Do(func() {
		DBInstance = &DBConnection{
			conn: "This is Mysql Connection sting ",
		}
	})

	return DBInstance
}

func main() {
	dbConnection := NewDBConnection()
	fmt.Println(dbConnection.conn)
}
