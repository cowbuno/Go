package main

type Customer struct {
	name  string
	email string
}

func newCustomer(name string, email string) *Customer {
	return &Customer{
		name:  name,
		email: email,
	}
}

func (c *Customer) update() {}

func getEmail(c *Customer) string {
	return c.email
}
