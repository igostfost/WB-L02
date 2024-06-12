package main

/*
Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
*/
import "fmt"

// Command интерфейс
type Command interface {
	execute()
}

// Light получатель в виде класса света
type Light struct {
	isOn bool
}

func (l *Light) On() {
	fmt.Println("light on")
	l.isOn = true
}

func (l *Light) Off() {
	fmt.Println("light off")
	l.isOn = false
}

// Реализуем конкретные команды
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) execute() {
	c.light.Off()
}

// RemoteControl является инициатором
type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.command = command
}

func (rc *RemoteControl) PressButton() {
	rc.command.execute()
}

// Клиентский код
func main() {
	light := &Light{}

	lightOnCommand := &LightOnCommand{light: light}
	lightOffCommand := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	//Включаем свет
	remote.SetCommand(lightOnCommand)
	remote.PressButton()

	//Выключаем свет
	remote.SetCommand(lightOffCommand)
	remote.PressButton()
}
