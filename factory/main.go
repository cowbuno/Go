package main

import "fmt"

type IGun interface {
	getCharachteristic() string
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) getCharachteristic() string {
	return fmt.Sprintf("Gun name:%s \nGun power:%d \n", g.name, g.power)
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	fmt.Printf(ak47.getCharachteristic())
	fmt.Println()
	fmt.Printf(musket.getCharachteristic())

}
