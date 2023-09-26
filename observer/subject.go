package observer

type Subject interface {
	addObserver(o Observer)
	deleteObserver(o Observer)
	notifyObservers()
	printAllObserver()
}
