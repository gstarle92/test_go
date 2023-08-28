package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Name Of person is :", p.Name)
	fmt.Println("Name Of person is :", p.Age)
}
func NewGreet(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}

}
func main() {

	person := NewGreet("gokul", 30)
	fmt.Println("Name Of person is :", person)
}
