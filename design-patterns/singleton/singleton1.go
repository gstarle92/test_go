package main

type singleton map[string]string

var (
	instance singleton
)

func NewSingleton() singleton {

	if instance == nil {
		instance = make(singleton)
	}
	return instance
}
func main() {

	NewSingleton()
}
