package observer

type Observer interface {
	update(string)
	getEmail() string
	printing()
}
