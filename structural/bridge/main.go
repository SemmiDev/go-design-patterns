package main

import (
	"log"
)

type BuildAPI interface {
	BuildCar(string)
}

type SmallCar struct{}

func (r *SmallCar) BuildCar(name string) {
	log.Println("build: ", name)
}

type MediumCar struct{}

func (r *MediumCar) BuildCar(name string) {
	log.Println("build: ", name)
}

type Design interface {
	Build()
}

type Car struct {
	name  string
	build BuildAPI
}

func (c *Car) Build() {
	c.build.BuildCar(c.name)
}

func main() {
	smallCar := &Car{"small car name", &SmallCar{}}
	mediumCar := &Car{"medium car name", &MediumCar{}}

	smallCar.Build()
	mediumCar.Build()
}
