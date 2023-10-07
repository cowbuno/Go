package decorator

import "fmt"

type TeleportCharacter struct {
	Character
}

func NewTeleportCharacter(c *BasicCharacter) *TeleportCharacter {
	c.SetDamge(c.Attack() - 2)
	return &TeleportCharacter{
		Character: c,
	}
}

func (c *TeleportCharacter) Teleport() {
	fmt.Println(c.GetName(), "character is teleported")
}
