package main

import "fmt"

func main() {
	airEezy := NewItem("air eezy")

	bob := NewCustomer("bob@gmail.com")
	anny := NewCustomer("anny@gmail.com")
	steve := NewCustomer("steve@gmail.com")

	airEezy.Regiter(bob)
	airEezy.Regiter(anny)
	airEezy.Regiter(steve)

	airEezy.updateAvailability()
}

type Subject interface {
	Regiter(o Observer)
	Unregister(o Observer)
	NotifyAll()
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) Regiter(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Unregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) NotifyAll() {
	for _, obs := range i.observerList {
		obs.Update(i.name)
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()

}

func removeFromSlice(obsList []Observer, obsToRemove Observer) []Observer {
	for i, observer := range obsList {
		if observer.GetId() == obsToRemove.GetId() {
			lengthObsList := len(obsList)
			obsList[i] = obsList[lengthObsList-1]
			return obsList[:lengthObsList-1]
		}
	}
	return obsList
}

type Observer interface {
	Update(string)
	GetId() string
}

type Customer struct {
	id string
}

func NewCustomer(id string) *Customer {
	return &Customer{
		id: id,
	}
}

func (c *Customer) Update(item string) {
	fmt.Printf("Send it to %s that %s is in stock\n", c.id, item)
}

func (c *Customer) GetId() string {
	return c.id
}
