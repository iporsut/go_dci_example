package model

type Accounter interface {
	SetName(name string)
	Name() string
	SetBalance(balance float32)
	Balance() float32
}

type Account struct {
	name    string
	balance float32
}

func NewAccount(name string, balance float32) *Account {
	return &Account{name, balance}
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) SetBalance(balance float32) {
	a.balance = balance
}

func (a *Account) Balance() float32 {
	return a.balance
}
