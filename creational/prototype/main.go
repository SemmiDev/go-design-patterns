package main

import "fmt"

type Position string

const (
	Manager                 Position = "Manager"
	SeniorFrontEndDeveloper Position = "Senior Front End Developer"
)

type Employee struct {
	Name     string
	Position Position
	Sallary  float64
}

func NewEmployeeManager(name string) *Employee {
	return &Employee{
		Name:     name,
		Position: Manager,
		Sallary:  30_000_000,
	}
}

func NewEmployeeSeniorFE(name string) *Employee {
	return &Employee{
		Name:     name,
		Position: SeniorFrontEndDeveloper,
		Sallary:  50_000_000,
	}
}

func main() {
	sam := NewEmployeeManager("sammi")
	izzah := NewEmployeeManager("izzah")
	dev := NewEmployeeSeniorFE("sam")
	fmt.Println(sam)
	fmt.Println(izzah)
	fmt.Println(dev)
}
