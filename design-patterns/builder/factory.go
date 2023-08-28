package main

import "fmt"

type FanI interface {
	GetCompany() string
	GetVoltage() string
}
type Fan struct {
	companyName string
	voltage     string
}

func (f *Fan) GetCompany() string {
	return f.companyName
}

func (f *Fan) GetVoltage() string {
	return f.voltage
}

type cromptonFans struct {
	Fan
}

type bajajFans struct {
	Fan
}

func FanFactory(FanType string) FanI {
	if FanType == "Crompton" {
		return &cromptonFans{
			Fan{
				companyName: "crompton",
				voltage:     "12 volt",
			}}

	}
	if FanType == "Bajaj" {
		return &bajajFans{Fan{
			companyName: "Bajaj",
			voltage:     "10 volt",
		}}
	}
	return nil
}

func main() {
	fancrompton := FanFactory("Crompton")

	fmt.Println("Fan Comapany name", fancrompton.GetCompany())
	fmt.Println("Fan Comapany name", fancrompton.GetVoltage())

	fanbajaj := FanFactory("Bajaj")

	fmt.Println("Fan Comapany name", fanbajaj.GetCompany())
	fmt.Println("Fan Comapany name", fanbajaj.GetVoltage())
}
