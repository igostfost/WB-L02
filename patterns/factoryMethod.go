package main

/*
Фабричный метод — это порождающий паттерн проектирования,
который решает проблему создания различных продуктов, без указания конкретных классов продуктов.
*/

/*
Сперва, мы создадим интерфейс ICar, который определяет все методы будущих машин.
Также имеем структуру Car, которая реализует интерфейс ICar.
Две конкретные машины — Lada и Mercedes — обе включают в себя структуру Car и не напрямую реализуют все методы от ICar.

CarFactory служит фабрикой, которая создает машину нужного типа в зависимости от аргумента на входе.
Клиентом служит main.go .
Вместо прямого взаимодействия с объектами Lada или Mercedes, он создает экземпляры разных машин
при помощи CarFactory, используя для контроля изготовления только параметры в виде строк.
*/
import "fmt"

type ICar interface {
	setName(name string)
	SetSpeed(speed int)
	GetSpeed() int
	GetName() string
}

type Car struct {
	name  string
	speed int
}

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) SetSpeed(speed int) {
	c.speed = speed
}

func (c *Car) GetSpeed() int {
	return c.speed
}

func (c *Car) GetName() string {
	return c.name
}

type Lada struct {
	Car
}

func NewLada() ICar {
	return &Lada{
		Car{
			name:  "Lada",
			speed: 100},
	}
}

type Mercedes struct {
	Car
}

func NewMercedes() ICar {
	return &Mercedes{
		Car{
			name:  "Mercedes",
			speed: 240},
	}
}

func CarFactory(carType string) (ICar, error) {
	if carType == "Lada" {
		return NewLada(), nil
	}
	if carType == "Mercedes" {
		return NewMercedes(), nil
	}
	return nil, fmt.Errorf("Wrong car type")
}

func main() {
	lada, _ := CarFactory("Lada")
	Mercedes, _ := CarFactory("Mercedes")

	fmt.Printf("Car: %s\nSpeed: %d\n", lada.GetName(), lada.GetSpeed())
	fmt.Println()
	fmt.Printf("Car: %s\nSpeed: %d\n", Mercedes.GetName(), Mercedes.GetSpeed())
}
