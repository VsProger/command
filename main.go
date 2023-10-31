package main

import "fmt"

type Command interface {
	Execute()
}

type Device interface {
	TurnOn()
	TurnOff()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) TurnOn() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}

func (t *Tv) TurnOff() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.TurnOn()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.TurnOff()
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.Press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.Press()
}
