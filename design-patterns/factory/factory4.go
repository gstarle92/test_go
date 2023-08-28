package main

import "fmt"

type Laptop interface {
	GetLaptopName() string
}

type HP struct {
	Name string
}

type Lenovo struct {
	Name string
}

func (hp *HP) GetLaptopName() string {
	return hp.Name
}

func (l *Lenovo) GetLaptopName() string {
	return l.Name
}

func LaptpFactory(brand string) (Laptop, error) {
	if brand == "HP" {
		return &HP{Name: "HP"}, nil
	}

	if brand == "Lenovo" {
		return &Lenovo{Name: "Lenovo"}, nil
	}
	return nil, fmt.Errorf("Invalid Brand Name")
}

func main() {

	laptopHP, _ := LaptpFactory("HP")
	fmt.Println(laptopHP.GetLaptopName())

	lenovo, _ := LaptpFactory("Lenovo")
	fmt.Println(lenovo.GetLaptopName())
}
