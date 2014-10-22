package context

import (
	"github.com/iporsut/go_dci_example/model"
	"github.com/iporsut/go_dci_example/role"
)

type MoneyTransfer struct {
	source, destination model.Accounter
}

func NewMoneyTransfer(source, destination model.Accounter) *MoneyTransfer {
	return &MoneyTransfer{&role.TransferSource{source}, destination}
}

func (mt *MoneyTransfer) Execute(amount float32) {
	mt.source.(*role.TransferSource).TransferTo(mt.destination, amount)
}
