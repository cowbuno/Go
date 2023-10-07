package decorator

type Character interface {
	Attack() int
	GetHealth() int
	TakeDamage(int)
	GetName() string
	SetDamge(int)
}
