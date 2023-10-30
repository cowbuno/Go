package command

import "fmt"

func (l *Light) SwitchOn() {
	l.isOn = true
	fmt.Println("Light is ON")
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(l *Light) *LightOnCommand {
	return &LightOnCommand{
		light: l,
	}
}

func (loc *LightOnCommand) Execute() {
	loc.light.SwitchOn()
}
