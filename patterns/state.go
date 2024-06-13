package main

import "fmt"

/*
Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.

Поведения, зависящие от состояния, переезжают в отдельные классы.
Первоначальный класс хранит ссылку на один из таких объектов-состояний и делегирует ему работу.
*/

// State - интерфейс состояния
type State interface {
	insertCoin()
	pressButton()
	dispenseCoffee()
}

type CoffeeMachine struct {
	state State
}

func (cm *CoffeeMachine) setState(state State) {
	cm.state = state
}

func (cm *CoffeeMachine) insertCoin() {
	cm.state.insertCoin()
}

func (cm *CoffeeMachine) pressButton() {
	cm.state.pressButton()
}

func (cm *CoffeeMachine) dispenseCoffee() {
	cm.state.dispenseCoffee()
}

// Реализация состояний

// NoCoinState - состояние "Нет монет"
type NoCoinState struct {
	coffeeMachine *CoffeeMachine
}

func NewNoCoinState(coffeeMachine *CoffeeMachine) *NoCoinState {
	return &NoCoinState{coffeeMachine: coffeeMachine}
}

func (s *NoCoinState) insertCoin() {
	fmt.Println("Монета вставлена")
	s.coffeeMachine.setState(NewHasCoinState(s.coffeeMachine))
}

func (s *NoCoinState) pressButton() {
	fmt.Println("Сначала вставьте монету")
}

func (s *NoCoinState) dispenseCoffee() {
	fmt.Println("Сначала вставьте монету")
}

// HasCoinState - состояние "Монета вставлена"
type HasCoinState struct {
	coffeeMachine *CoffeeMachine
}

func NewHasCoinState(coffeeMachine *CoffeeMachine) *HasCoinState {
	return &HasCoinState{coffeeMachine: coffeeMachine}
}

func (s *HasCoinState) insertCoin() {
	fmt.Println("Монета уже вставлена")
}

func (s *HasCoinState) pressButton() {
	fmt.Println("Приготовление кофе")
	s.coffeeMachine.setState(NewMakingCoffeeState(s.coffeeMachine))
}

func (s *HasCoinState) dispenseCoffee() {
	fmt.Println("Кофе еще не готов")
}

// MakingCoffeeState - состояние "Кофе готовится"
type MakingCoffeeState struct {
	coffeeMachine *CoffeeMachine
}

func NewMakingCoffeeState(coffeeMachine *CoffeeMachine) *MakingCoffeeState {
	return &MakingCoffeeState{coffeeMachine: coffeeMachine}
}

func (s *MakingCoffeeState) insertCoin() {
	fmt.Println("Кофе уже готовится, подождите")
}

func (s *MakingCoffeeState) pressButton() {
	fmt.Println("Кофе уже готовится")
}

func (s *MakingCoffeeState) dispenseCoffee() {
	fmt.Println("Ваш кофе готов!")
	s.coffeeMachine.setState(NewNoCoinState(s.coffeeMachine))
}

// клиентский код
func main() {
	coffeeMachine := &CoffeeMachine{state: NewNoCoinState(nil)}
	coffeeMachine.setState(NewNoCoinState(coffeeMachine))

	coffeeMachine.insertCoin()     // Монета вставлена
	coffeeMachine.pressButton()    // Приготовление кофе
	coffeeMachine.dispenseCoffee() // Ваш кофе готов!

	coffeeMachine.pressButton()    // Сначала вставьте монету
	coffeeMachine.insertCoin()     // Монета вставлена
	coffeeMachine.insertCoin()     // Монета уже вставлена
	coffeeMachine.pressButton()    // Приготовление кофе
	coffeeMachine.dispenseCoffee() // Ваш кофе готов!
}
