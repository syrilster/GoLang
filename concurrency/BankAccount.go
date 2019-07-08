package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type Account struct {
	sync.Mutex
	balance int
	open    bool
}

func Open(amount int) (*Account, error) {
	if amount < 0 {
		return nil, errors.New("Can't open account with negative balance deposit")
	}
	return &Account{balance: amount, open: true}, nil
}

func (account *Account) Close() (int, bool) {
	account.Lock()
	defer account.Unlock()
	if !account.open {
		return 0, false
	}
	account.open = false
	return account.balance, true
}

func (account *Account) Deposit(amount int) (bool, error) {
	defer account.Unlock()
	account.Lock()
	if !account.open || account.balance+amount < 0 {
		return false, errors.New("Account is in Invalid status")
	}
	fmt.Printf("Customer depositing an amount of %d into account \n", amount)
	account.balance += amount
	return true, nil
}

func (account *Account) Balance() int {
	return account.balance
}

func (account *Account) Withdraw(amount int) (bool, error) {
	defer account.Unlock()
	account.Lock()
	if amount > account.balance {
		return false, errors.New("Insufficient Balance. Withdrawal Failed for amount !!")
	}

	fmt.Printf("Customer withdrawing an amount of %d \n", amount)
	account.balance -= amount
	return true, nil
}

func displayError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println("Go Bank Account example")
	var accountStatus = "Closed"
	account, err := Open(5000)
	displayError(err)
	fmt.Printf("Account 'a' opened with initial balance of %d \n", account.balance)
	//account.Close()
	_, err = account.Deposit(200)

	displayError(err)
	_, err = account.Withdraw(4000)

	displayError(err)
	if account.open {
		accountStatus = "Open"
	}
	fmt.Printf("Current Account Balance is %d and account status is %s", account.Balance(), accountStatus)
}
