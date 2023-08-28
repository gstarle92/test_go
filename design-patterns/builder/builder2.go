package main

import "fmt"

type BLaptop interface {
	GetLaptopName() string
}

type Delllaptop struct {
	laptopName string
}
type HPLaptop struct {
	laptopName string
}

func (d *Delllaptop) GetLaptopName() string {
	return d.laptopName
}

func (h *HPLaptop) GetLaptopName() string {
	return h.laptopName
}

func NewDellLaptop() BLaptop {
	return &Delllaptop{
		laptopName: "This is Dell Laptop ",
	}
}

func NewHPLaptop() BLaptop {
	return &HPLaptop{
		laptopName: "This is HP Laptop ",
	}
}

func LaptopBuilder(typeLaptop string) BLaptop {
	if typeLaptop == "Dell" {
		return NewDellLaptop()
	}
	if typeLaptop == "HP" {
		return NewHPLaptop()
	}
	return nil
}

func main() {

	dellLaptop := LaptopBuilder("Dell")
	fmt.Println(dellLaptop.GetLaptopName())

	hpLaptop := LaptopBuilder("HP")
	fmt.Println(hpLaptop.GetLaptopName())
}
