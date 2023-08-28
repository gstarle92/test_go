package main

import "fmt"

type LaptopI interface {
	SetRAM()
	SetCPU()
	SetMotherboard()
	GetLaptop() Laptop
}
type Dell struct {
	ram string
	cpu string
	mb  string
}

type Laptop struct {
	ram string
	cpu string
	mb  string
}

func newDell() *Dell {
	return &Dell{}

}

func (d *Dell) SetRAM() {
	d.ram = "8 GB Ram "
}
func (d *Dell) SetCPU() {
	d.cpu = "inter icore  "
}
func (d *Dell) SetMotherboard() {
	d.mb = "Asus "
}

func (d *Dell) GetLaptop() Laptop {
	return Laptop{
		ram: d.ram,
		cpu: d.cpu,
		mb:  d.mb,
	}
}
func GetLaptopBuilder(laptopType string) LaptopI {
	if laptopType == "Dell" {
		return newDell()
	}
	return nil
}

type builder struct {
	builder LaptopI
}

func newLaptopBuilder(l LaptopI) *builder {
	return &builder{
		builder: l,
	}
}

func (b *builder) BuildLaptop() Laptop {
	b.builder.SetCPU()
	b.builder.SetRAM()
	b.builder.SetMotherboard()
	return b.builder.GetLaptop()
}
func main() {

	dell := GetLaptopBuilder("Dell")
	lbuilder := newLaptopBuilder(dell)
	laptop := lbuilder.BuildLaptop()

	fmt.Println(laptop.ram)
	fmt.Println(laptop.cpu)
	fmt.Println(laptop.mb)
}
