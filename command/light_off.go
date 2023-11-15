package command

import "fmt"

func (l *Light) SwitchOff() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(l *Light) *LightOffCommand {
	return &LightOffCommand{
		light: l,
	}
}
func (loc *LightOffCommand) Execute() {
	loc.light.SwitchOff()
}