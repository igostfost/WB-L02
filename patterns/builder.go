package main

/*
Паттерн строитель используется для конструирования сложных объктов пошагово. В Этом примере
мы строим последовательно дом, установив двери, окна и этажи.
*/
import "fmt"

type House struct {
	doorType   string
	windowType string
	floorCount int
}

type Builder interface {
	SetDoor()
	SetWindow()
	SetFloor()
}

type normalBuilder struct {
	doorType   string
	windowType string
	floorCount int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (nb *normalBuilder) SetDoor() {
	nb.doorType = "normalDoor"
}

func (nb *normalBuilder) SetWindow() {
	nb.windowType = "normalWindow"
}
func (nb *normalBuilder) SetFloor() {
	nb.floorCount = 2
}

func (nb *normalBuilder) getHouse() House {
	fmt.Println("BuildHouse")
	return House{
		doorType:   nb.doorType,
		windowType: nb.windowType,
		floorCount: nb.floorCount,
	}
}

func (nb *normalBuilder) BuildHouse() House {
	nb.SetDoor()
	nb.SetWindow()
	nb.SetFloor()
	return nb.getHouse()
}

func main() {
	builder := newNormalBuilder()
	newHouse := builder.BuildHouse()

	fmt.Printf("doorType in house: %s\n", newHouse.doorType)
	fmt.Printf("windowType in house: %s\n", newHouse.windowType)
	fmt.Printf("floorCount in house: %d\n", newHouse.floorCount)
}
