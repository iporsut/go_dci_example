package main

import (
	"fmt"
	"github.com/iporsut/go_dci_example/context"
	"github.com/iporsut/go_dci_example/model"
)

func main() {
	source := model.NewAccount("Weerasak", 1000)
	dest := model.NewAccount("Kanokon", 1000)
	moneyTransfer := context.NewMoneyTransfer(source, dest)

	moneyTransfer.Execute(500)

	fmt.Println("Weerasak balance ", source.Balance())
	fmt.Println("Kanokon balance ", dest.Balance())
}
