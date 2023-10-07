package decorator

import "fmt"

type BasicCharacter struct {
	name   string
	health int
	damage int
}

func NewBasicCharacter(name string) *BasicCharacter {
	return &BasicCharacter{
		name:   name,
		health: 100,
		damage: 5,
	}

}

func (c *BasicCharacter) Attack() int {
	return c.damage
}

func (c *BasicCharacter) SetDamge(damage int) {
	c.damage = damage
}

func (c *BasicCharacter) TakeDamage(amount int) {
	c.health -= amount
	if c.health <= 0 {
		fmt.Println(c.name, "is dead")
	}
}

func (c *BasicCharacter) GetHealth() int {
	return c.health
}

func (c *BasicCharacter) GetName() string {
	return c.name
}
