package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("cannot withdraw")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner:owner, balance:0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Return balance of your account
func (a Account) Balance() int{
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//Change owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(),"'s account.\nBalance: ",a.Balance(),"$")
}
