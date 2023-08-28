package main

import "fmt"

type Student struct {
	Roll  int
	FName string
	LName string
}

//Factory method is a creational design pattern which solves the problem of creating product objects without
//specifying their concrete classes.
//The Factory Method defines a method, which should be used for creating objects instead of using a direct constructor call ( new operator).

func (s Student) getStudentDetails() string {

}

func NewStudent(roll int, fname string, lname string) Student {
	return Student{
		Roll:  roll,
		FName: fname,
		LName: lname,
	}
}

func main() {

	fmt.Println(student)
}
