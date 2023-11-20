package main

import "fmt"

type Button struct {
	Command Command
}

func (b *Button) PressButton() {
	b.Command.Execute()
}

type Command interface {
	Execute()
}

type OnCommand struct {
	Device Device
}

func (oc *OnCommand) Execute() {
	oc.Device.On()
}

type OffCommand struct {
	Device Device
}

func (oc *OffCommand) Execute() {
	oc.Device.Off()
}

type Device interface {
	Off()
	On()
}

type TV struct {
	IsRunning bool
}

func (t *TV) Off() {
	t.IsRunning = false
	fmt.Println("tv is turn of")
}

func (t *TV) On() {
	t.IsRunning = true
	fmt.Println("tv is turn on")
}

func main() {
	tv := &TV{}

	OnCommand := OnCommand{
		Device: tv,
	}

	offCommand := OffCommand{
		Device: tv,
	}

	OnButton := Button{
		Command: &OnCommand,
	}

	OffButton := Button{
		Command: &offCommand,
	}

	OnButton.PressButton()
	OffButton.PressButton()
}
