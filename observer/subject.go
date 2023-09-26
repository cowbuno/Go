package observer

type Subject interface {
	AddObserver(o Observer)
	DeleteObserver(o Observer)
	ItemIsAvailable()
	PrintAllObserver()
}
