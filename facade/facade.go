package main

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU is freeze")
}

func (c *CPU) Jump(position int) {
	fmt.Printf("SPU jump to %d \n", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU execute")
}

type Memory struct{}

func (m *Memory) Load(position int, data string) {
	fmt.Printf("Memory load data %d to position %s\n", position, data)
}

type HardDrive struct{}

func (hd *HardDrive) Read(position, size int) string {
	data := fmt.Sprintf("hard drive read data from position %d with size %d", position, size)
	fmt.Println(data)
	return data
}

type ComputerFacade struct {
	CPU       *CPU
	Memory    *Memory
	HardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		CPU:       &CPU{},
		Memory:    &Memory{},
		HardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	c.CPU.Freeze()
	c.Memory.Load(0, "boot_loader")
	c.CPU.Jump(0)
	c.CPU.Execute()
}

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
