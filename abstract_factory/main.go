package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 1,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 2,
		},
	}
}

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 9,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 8,
		},
	}
}

type IShoe interface {
	GetCharachteristic() string
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) GetCharachteristic() string {
	return fmt.Sprintf("Shoe logo :%s \nShoe size:%v\n", s.logo, s.size)
}

type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

type IShirt interface {
	GetCharachteristic() string
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) GetCharachteristic() string {
	return fmt.Sprintf("Shirt logo :%s \nShoe size:%v\n ", s.logo, s.size)
}

type AdidasShirt struct {
	Shirt
}

type NikeShirt struct {
	Shirt
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	fmt.Println(nikeShoe.GetCharachteristic())
	fmt.Println(nikeShirt.GetCharachteristic())
	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	fmt.Println(adidasShoe.GetCharachteristic())
	fmt.Println(adidasShirt.GetCharachteristic())

}
