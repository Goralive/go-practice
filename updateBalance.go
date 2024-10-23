package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cust *customer, transaction transaction) error {
	var balance float64 = (*cust).balance
	if transaction.transactionType == transactionDeposit {
		*&cust.balance += transaction.amount
		return nil
	}

	if transaction.transactionType == transactionWithdrawal {
		if balance-transaction.amount < 0 {
			return errors.New("insufficient funds")
		} else {
			*&cust.balance -= transaction.amount
			return nil
		}
	}
	return errors.New("unknown transaction type")

}
