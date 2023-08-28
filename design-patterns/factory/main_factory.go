package main

import "fmt"

type CarI interface {
	GetCarModel() string
	GetCarType() string
	SetCarType(carType string)
	SetCarModel(carModel string)
}

type car struct {
	carType  string
	carModel string
}

func (c *car) SetCarType(carType string) {
	c.carType = carType
}

func (c *car) SetCarModel(carModel string) {
	c.carModel = carModel
}

func (c *car) GetCarModel() string {
	return c.carType
}

func (c *car) GetCarType() string {
	return c.carModel
}

type Maruti800 struct {
	car
}

func getCar(carType string) (CarI, error) {
	if carType == "Maruti800" {
		return NewMaruti800(), nil
	}

	return nil, fmt.Errorf("Wrong Car type passed")
}
func NewMaruti800() CarI {
	return &Maruti800{
		car: car{
			carType:  "small Car",
			carModel: "14454dsajhju",
		},
	}
}

func main() {
	car, _ := getCar("Maruti800")

	printDetails(car)
}

func printDetails(c CarI) {
	fmt.Printf("Car Model: %s", c.GetCarModel())
	fmt.Println()
	fmt.Printf("Car Type: %s", c.GetCarType())
	fmt.Println()

}
