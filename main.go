package main

import (
	"fmt"
	"github.com/iporsut/go_dci_example/context"
	"github.com/iporsut/go_dci_example/model"
)

func main() {
	source := model.NewAccount("Weerasak", 1000)
	dest := model.NewAccount("Kanokon", 1000)

	context.TransferMoney(source, dest, 500.0)

	fmt.Println("Weerasak balance ", source.Balance())
	fmt.Println("Kanokon balance ", dest.Balance())
}
