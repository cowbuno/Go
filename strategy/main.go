package main

import "fmt"

type DialogReciter interface {
	Recite()
}

type Superman struct{}

func (s *Superman) Recite() {
	fmt.Println("There is a superhero in all of us, ")
}

type Batman struct{}

func (b *Batman) Recite() {
	fmt.Println("It's not who I am underneath")
}

type Spiderman struct{}

func (s *Spiderman) Recite() {
	fmt.Println("Friendly neighbor spider")
}

type toy struct {
	dialogReciter DialogReciter
}

func NewToy(dr DialogReciter) *toy {
	return &toy{
		dialogReciter: dr,
	}
}

func (t *toy) Dialog() {
	t.dialogReciter.Recite()
}

func (t *toy) SetSuperHero(dr DialogReciter) {
	t.dialogReciter = dr
}

func main() {
	toy := NewToy(&Spiderman{})
	toy.Dialog()
	toy.SetSuperHero(&Batman{})
	toy.Dialog()
	toy.SetSuperHero(&Spiderman{})
	toy.Dialog()
}
