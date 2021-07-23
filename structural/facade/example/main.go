package main

import (
	"fmt"
)

type CustomerServiceProvider string

const (
	Dana  CustomerServiceProvider = "DANA"
	GoPay CustomerServiceProvider = "GO-PAY"
	OVO   CustomerServiceProvider = "OVO"
)

type Complain struct {
	CustomerEmail string
	Topic         string
	Desc          string
	Provider      CustomerServiceProvider
}

type Response struct {
	Msg string
}

func HandleComplain(c *Complain) Response {
	switch c.Provider {
	case Dana:
		resp := DanaComplainServiceProvider(c)
		return Response{Msg: resp}
	case GoPay:
		resp := GoPayComplainServiceProvider(c)
		return Response{Msg: resp}
	case OVO:
		resp := OVOComplainServiceProvider(c)
		return Response{Msg: resp}
	}

	return Response{
		Msg: "something went wrong",
	}
}

func main() {
	csp := HandleComplain(&Complain{
		CustomerEmail: "sammi@mail.com",
		Topic:         "XX",
		Desc:          "XX",
		Provider:      Dana,
	})

	fmt.Println(csp.Msg)
}
