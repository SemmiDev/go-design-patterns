package main

import "fmt"

type EWalletProviderTransfer interface {
	Transfer() *Payload
}

type DanaProvider struct {
	payload *Payload
}
type OVOProvider struct {
	payload *Payload
}
type LinkAjaProvider struct {
	payload *Payload
}

func (d *DanaProvider) Transfer() *Payload {
	return d.payload
}
func (d *OVOProvider) Transfer() *Payload {
	return d.payload
}
func (d *LinkAjaProvider) Transfer() *Payload {
	return d.payload
}

type EWalletType int

const (
	Dana EWalletType = 1 << iota
	OVO
	LinkAja
)

type Payload struct {
	fromAccountProvider EWalletType
	toAccountProvider   EWalletType
	amount              uint64
}

func EWalletProviderTransferFactory(t EWalletType, payload *Payload) EWalletProviderTransfer {
	payload.fromAccountProvider = t
	switch t {
	case Dana:
		return &DanaProvider{payload: payload}
	case OVO:
		return &OVOProvider{payload: payload}
	case LinkAja:
		return &LinkAjaProvider{payload: payload}
	default:
		return nil
	}
}

func main() {
	tf1 := EWalletProviderTransferFactory(Dana, &Payload{
		toAccountProvider: OVO,
		amount:            200000,
	})

	tf2 := EWalletProviderTransferFactory(LinkAja, &Payload{
		toAccountProvider: OVO,
		amount:            500000,
	})

	fmt.Println(tf1.Transfer())
	fmt.Println(tf2.Transfer())
}
