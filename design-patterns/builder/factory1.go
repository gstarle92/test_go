package main

import "fmt"

type LaptopInterface interface {
	GetRam() string
	GetCPU() string
}
type Dell struct {
	ram string
	cpu string
}

func (l *Dell) GetRam() string {
	return l.ram
}

func (l *Dell) GetCPU() string {
	return l.cpu
}

type HP struct {
	ram string
	cpu string
}

func (l *HP) GetRam() string {
	return l.ram
}

func (l *HP) GetCPU() string {
	return l.cpu
}

func LaptopFactory(laptoptype string) LaptopInterface {
	if laptoptype == "Dell" {
		return &Dell{
			ram: "8 Gb Ram ",
			cpu: "Intel icore ",
		}
	}

	if laptoptype == "HP" {
		return &HP{
			ram: "8 Gb Ram ",
			cpu: "Intel icore ",
		}
	}

	return nil
}

func main() {
	dell := LaptopFactory("Dell")

	fmt.Println("Dell LaptopRam", dell.GetCPU())
	fmt.Println("Dell LaptopRam", dell.GetRam())

	hp := LaptopFactory("HP")

	fmt.Println("hp LaptopRam", hp.GetCPU())
	fmt.Println("hp LaptopRam", hp.GetRam())
}
