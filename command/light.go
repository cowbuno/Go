package command

type Light struct {
	isOn bool
}

func NewLight() *Light {
	return &Light{}
}
