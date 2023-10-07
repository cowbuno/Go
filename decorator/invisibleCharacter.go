package decorator

import "fmt"

type InvisibleCharacter struct {
	Character
}

func NewInvisibleCharacter(c *BasicCharacter) *InvisibleCharacter {
	c.SetDamge(c.Attack() - 4)
	return &InvisibleCharacter{
		Character: c,
	}
}

func (c *InvisibleCharacter) BeInvisible() {
	fmt.Println(c.GetName(), "character is invisible")
}

func (c *InvisibleCharacter) BeVisible() {
	fmt.Println(c.GetName(), "character is visible")
}
