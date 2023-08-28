package main

import "fmt"

type FanI interface {
	GetCompany() string
	GetVoltage() string
}

type Fan struct {
	comapnyName string
	voltage     string
}

func (f *Fan) GetCompany() string {
	return f.comapnyName
}

func (f *Fan) GetVoltage() string {
	return f.voltage
}

type Crompton struct {
	Fan
}
type Havells struct {
	Fan
}

func NewCromptonFan() FanI {
	return &Crompton{
		Fan: Fan{
			comapnyName: "Crompton",
			voltage:     "12V",
		},
	}
}

func NewHevellsFan() FanI {
	return &Havells{
		Fan: Fan{
			comapnyName: "Hevells",
			voltage:     "12V",
		},
	}
}

func FanFactory(brand string) (FanI, error) {

	if brand == "Crompton" {
		return NewCromptonFan(), nil
	}
	if brand == "Hevells" {
		return NewHevellsFan(), nil
	}
	return nil, fmt.Errorf("invalid Fan Brand")

}

func main() {
	cromptonFan, err := FanFactory("Crompton")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Fan Campany:", cromptonFan.GetCompany())
	fmt.Println("Fan Vontage:", cromptonFan.GetVoltage())

	hevellsFan, err := FanFactory("Hevells")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Fan Campany:", hevellsFan.GetCompany())
	fmt.Println("Fan Vontage:", hevellsFan.GetVoltage())

}
