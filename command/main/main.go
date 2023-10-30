package main

import "SDP/command"

func main() {
	// Client Code

	light := command.NewLight()
	lightOn := command.NewLightOnCommand(light)
	lightOff := command.NewLightOffCommand(light)

	remote := command.RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()
}
