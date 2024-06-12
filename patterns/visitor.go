package main

/*
Посетитель — это поведенческий паттерн проектирования,
который позволяет добавлять в программу новые операции, не изменяя классы объектов,
над которыми эти операции могут выполняться.

*/
import "fmt"

// Visitor интерфейс нашего посетителя
type Visitor interface {
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// AreaCalculate Посетитель, который считает площадь фигур
type AreaCalculate struct {
	area float64
}

// При визите в круг - считаем площадь круга
func (ac *AreaCalculate) visitForCircle(c *Circle) {
	ac.area = 3.14 * float64(c.radius) * float64(c.radius)
}

// При визите в прямоуголник - считаем площадь прямоугольника
func (ac *AreaCalculate) visitForRectangle(r *Rectangle) {
	ac.area = r.a * r.b
}

// Shape создаем фигуры
type Shape interface {
	getType() string
	accept(Visitor)
}

type Rectangle struct {
	a, b float64
}

func (r *Rectangle) getType() string {
	return "rectangle"
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

type Circle struct {
	radius int
}

func (c *Circle) getType() string {
	return "Circle"
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func main() {
	newCircle := Circle{radius: 5}
	newRectangle := Rectangle{a: 3, b: 5}
	areaCalculator := &AreaCalculate{}

	newCircle.accept(areaCalculator)
	fmt.Printf("Circle area = %.2f\n", areaCalculator.area)

	newRectangle.accept(areaCalculator)
	fmt.Printf("Rectangle area = %.2f\n", areaCalculator.area)

}
