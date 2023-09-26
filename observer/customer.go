package observer

import "fmt"

type Customer struct {
	Name  string
	Email string
}

func NewCustomer(name string, email string) *Customer {
	return &Customer{
		Name:  name,
		Email: email,
	}
}

func (c *Customer) update(name string) {
	fmt.Println(name, " is an available now, come to the our shop")
}

func (c *Customer) getEmail() string {
	return c.Email
}
