package datamodels

import (
	"errors"
	"strconv"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

type Wallet struct {
	Currency string `bson:"currency" json:"currency"`
	Balance  int64  `bson:"balance" json:"balance"`
	IBAN     string `bson:"IBAN" json:"IBAN"`
}

func NewWallet(currency string, balance int64) *Wallet {
	wallet := new(Wallet)
	wallet.Currency = currency
	wallet.Balance = balance
	return wallet
}

func (wallet *Wallet) SetIBAN(bankCountryCode string, BIC string, accountNumber uint32) error {
	accountNumberToString := strconv.FormatUint(uint64(accountNumber), 10)
	iban := utils.ConcatenateStrings(bankCountryCode, BIC, accountNumberToString)
	if len(iban) != 24 {
		return errors.New("An error occured while setting wallets IBAN")
	}
	wallet.IBAN = iban
	return nil
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
