package main

import "fmt"

type IGun interface {
	GetPower() int
	SetPower(power int)
	GetName() string
	SetName(name string)
}

type Gun struct {
	power int
	name  string
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetPower() int {
	return g.power
}

func (g *Gun) GetName() string {
	return g.name
}

type Ak47 struct {
	Gun
}
type Mosket struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 Gun",
			power: 48,
		},
	}
}

func NewMosket() IGun {
	return &Mosket{
		Gun: Gun{
			name:  "Musket Gun",
			power: 8,
		},
	}
}

func GunFactory(factoryType string) (IGun, error) {
	if factoryType == "Ak47" {
		return NewAk47(), nil
	}

	if factoryType == "Musket" {
		return NewMosket(), nil
	}
	return nil, fmt.Errorf("invalid gun type")
}

func main() {
	ak47, _ := GunFactory("Ak47")
	fmt.Println(ak47.GetName())
	fmt.Println(ak47.GetPower())

	fmt.Println()
	musket, _ := GunFactory("Musket")
	fmt.Println(musket.GetName())
	fmt.Println(musket.GetPower())

}
