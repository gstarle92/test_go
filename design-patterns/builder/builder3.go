package main

import "fmt"

type LaptopI interface {
	SetRAM()
	SetCPU()
	SetMotherboard()
	GetLaptop() Laptop
}

func newLaptopBuilder(laptopType string) LaptopI {
	if laptopType == "Dell" {
		return newDellBuilder()
	}
	return nil
}

type Dell struct {
	ram          string
	cpu          string
	montherboard string
}

func newDellBuilder() *Dell {
	return &Dell{}
}

func (l *Dell) SetRAM() {
	l.ram = "8GM"
}

func (l *Dell) SetCPU() {
	l.cpu = "Inteli5"
}
func (l *Dell) SetMotherboard() {
	l.montherboard = "Asus motherboard"
}
func (l *Dell) GetLaptop() Laptop {
	return Laptop{
		ram:          l.ram,
		cpu:          l.cpu,
		montherboard: l.montherboard,
	}
}

type Laptop struct {
	ram          string
	cpu          string
	montherboard string
}

type Director struct {
	builder LaptopI
}

func newDirector(l LaptopI) *Director {
	return &Director{
		builder: l,
	}
}

func (d *Director) buildLaptop() Laptop {
	d.builder.SetRAM()
	d.builder.SetCPU()
	d.builder.SetMotherboard()
	return d.builder.GetLaptop()
}

func main() {
	dell := newLaptopBuilder("Dell")

	director := newDirector(dell)
	dellLaptop := director.buildLaptop()

	fmt.Printf("Normal House Door Type: %s\n", dellLaptop.ram)
	fmt.Printf("Normal House Window Type: %s\n", dellLaptop.cpu)
	fmt.Printf("Normal House Num type: %s\n", dellLaptop.montherboard)
}
