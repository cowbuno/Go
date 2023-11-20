package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningIntoComputer(comp Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	comp.InsertLightningPort()
}

type Computer interface {
	InsertLightningPort()
}

type Mac struct{}

func (m *Mac) InsertLightningPort() {
	fmt.Println("Lighting connector pludged to Mac")
}

type Windows struct{}

func (w *Windows) InserUSBPort() {
	fmt.Println("USB connector pludged to windows")
}

type WindowsAdapter struct {
	Windows *Windows
}

func (wa WindowsAdapter) InsertLightningPort() {
	fmt.Println("Adapter converts Lighting to USB")
	wa.Windows.InserUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningIntoComputer(mac)

	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		Windows: windows,
	}

	client.InsertLightningIntoComputer(windowsAdapter)
}
