package observer

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name:    name,
		inStock: false,
	}
}

func ItemIsAvailable(i *Item) {
	fmt.Println(i.name, " is come")
	i.inStock = true
	i.notifyObservers()
}

func (i *Item) notifyObservers() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}

}

func (i *Item) AddObserver(o Observer) {
	i.observerList = append(i.observerList, o)
}

func getIndex(observerList []Observer, email string) int {

	for i, obs := range observerList {
		current_email := obs.getEmail()
		if current_email == email {
			return i
		}
	}
	return -1
}

func (i *Item) DeleteObserver(o Observer) {
	deleted_email := o.getEmail()
	index := getIndex(i.observerList, deleted_email)
	i.observerList = append(i.observerList[:index], i.observerList[index+1:]...)

}

func (i *Item) PrintAllObserver() {
	for _, observer := range i.observerList {
		email := observer.getEmail()
		fmt.Println(email)
	}
}
