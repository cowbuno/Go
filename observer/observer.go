package main

type Observer interface {
	update(string)
	getEmail() string
	printing()
}
