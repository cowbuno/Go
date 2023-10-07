package decorator

import "fmt"

type FlyingCharacter struct {
	Character
}

func NewFlyingCharacter(c *BasicCharacter) *FlyingCharacter {
	c.SetDamge(c.Attack() - 1)
	return &FlyingCharacter{
		Character: c,
	}
}

func (c *FlyingCharacter) Fly() {
	fmt.Println(c.GetName(), "character is flying")
}
