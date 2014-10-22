package role

import (
	"github.com/iporsut/go_dci_example/model"
)

type TransferSource struct {
	model.Accounter
}

func (source TransferSource) TransferTo(destination model.Accounter, amount float32) {
	source.SetBalance(source.Balance() - amount)
	destination.SetBalance(destination.Balance() + amount)
}
