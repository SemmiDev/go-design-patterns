package main

import "fmt"

type Value *int

var Increment Value

func getInstance() Value {
	if Increment == nil {
		Increment = new(int)
	}
	return Increment
}

func IncrementValue(Value) {
	*Increment++
}

func GetValue() int {
	return *Increment
}

func main() {
	a := getInstance()
	b := getInstance()
	IncrementValue(a)
	IncrementValue(a)
	IncrementValue(b)
	IncrementValue(b)

	fmt.Println(a == b)
	fmt.Println(GetValue())
}
