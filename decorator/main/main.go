package main

import (
	"SDP/decorator"
	"fmt"
)

func main() {

	hero1 := decorator.NewBasicCharacter("hero 111")
	hero2 := decorator.NewBasicCharacter("hero 222")
	hero3 := decorator.NewBasicCharacter("hero 333")

	flyingHero := decorator.NewFlyingCharacter(hero1)
	invisibleHero := decorator.NewInvisibleCharacter(hero2)
	teleportHero := decorator.NewTeleportCharacter(hero3)

	fmt.Println(flyingHero.GetName())
	flyingHero.Fly()
	fmt.Println(flyingHero.Attack())

	fmt.Println()

	fmt.Println(invisibleHero.GetName())
	invisibleHero.BeInvisible()
	invisibleHero.BeVisible()
	fmt.Println(invisibleHero.Attack())

	fmt.Println()

	fmt.Println(teleportHero.GetName())
	teleportHero.Teleport()
	fmt.Println(teleportHero.Attack())

}
