package role

import (
	"github.com/iporsut/go_dci_example/model"
)

type transferSource struct {
	model.Accounter
}

func (source transferSource) TransferTo(destination model.Accounter, amount float32) {
	source.SetBalance(source.Balance() - amount)
	destination.SetBalance(destination.Balance() + amount)
}

func TransferSource(source model.Accounter) transferSource {
	return transferSource{source}
}
