package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Engineer's name: " + e.Name
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Manager's name: " + m.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := &Engineer{Name: "Michael"}
	PrintDetails(engineer)

	manager := &Manager{Name: "Stephen"}
	PrintDetails(manager)
}
