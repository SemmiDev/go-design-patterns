package main

import "fmt"

type Connection struct{}

var instance *Connection

func GetInstance() *Connection  {
	if instance == nil {
		instance = &Connection{}
	}
	return instance
}

func main() {
	var instance1 = GetInstance()
	var instance2 = GetInstance()
	fmt.Println(instance1 == instance2)
}