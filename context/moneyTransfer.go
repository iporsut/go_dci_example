package context

import (
	"github.com/iporsut/go_dci_example/model"
	"github.com/iporsut/go_dci_example/role"
)

func TransferMoney(source, destination model.Accounter, amount float32) {
	(&role.TransferSource{source}).TransferTo(destination, amount)
}

