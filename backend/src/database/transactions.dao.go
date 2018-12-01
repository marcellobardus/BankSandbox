package database

import (
	"errors"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"
)

func (connection *Connection) InsertRealisedTransaction(transaction *datamodels.Transaction) error {
	if transaction.TransactionHash == "" {
		return errors.New("Cannot insert a transaction with an undefined hash as realised")
	}
	err := database.C(RealisedTransactionsCollection).Insert(transaction)
	return err
}
