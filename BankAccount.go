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

func main() {
	fmt.Println("Go Bank Account example")
	//transactionChannel := make(chan bool)
	account, err := Open(5000)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Account 'a' opened with initial balance of %d \n", account.balance)
	}

	account.Deposit(200)
	_, err = account.Withdraw(40000)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Current Account Balance is: ", account.Balance())

}
