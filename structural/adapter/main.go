package main

import "fmt"

type TelegramLegacy struct{}

func (t *TelegramLegacy) SendMessageOldWay() string {
	return "send message telegram old way"
}

type MessageAdapter interface {
	SendMessage() string
}

type Adapter struct {
	telegramLegacy *TelegramLegacy
}

func TelegramAdapter() *Adapter {
	return &Adapter{&TelegramLegacy{}}
}

func (t *Adapter) SendMessage() string {
	fmt.Println("calling from adapter")
	return t.telegramLegacy.SendMessageOldWay()
}

func main() {
	telegram := TelegramAdapter()
	fmt.Println(telegram.SendMessage())
}
