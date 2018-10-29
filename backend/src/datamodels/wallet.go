package datamodels

import (
	"errors"
)

type Wallet struct {
	Currency string `bson:"currency" json:"currency"`
	Balance  int64  `bson:"balance" json:"balance"`
}

func NewWallet(currency string, balance int64) *Wallet {
	wallet := new(Wallet)
	wallet.Currency = currency
	wallet.Balance = balance
	return wallet
}

func (wallet *Wallet) IncreaseBalance(balance int64) error {
	if balance <= 0 {
		return errors.New("Balance cannot be less than 0 or be equal to 0")
	}
	wallet.Balance += balance
	return nil
}

func (wallet *Wallet) DecreaseBalance(balance int64) error {
	if balance <= 0 {
		return errors.New("Balance cannot be less than 0 or be equal to 0")
	}
	wallet.Balance -= balance
	return nil
}
