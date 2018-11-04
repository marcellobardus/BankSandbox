package datamodels

import (
	"log"
	"time"
)

type TransactionStatus int8

const (
	// Unconfirmed transaction cannot be send bcause its fee is undefined
	Unconfirmed TransactionStatus = 0
	// Sent => Transction is waiting for the nearest session
	Sent TransactionStatus = 1
	// Realised => Sender and Recipient wallets balances are updated
	Realised TransactionStatus = 2
	// Cancelled => Transaction is cancelled
	Cancelled TransactionStatus = 3
	// Frozen => Transaction is frozen
	Frozen TransactionStatus = 4
)

type Transaction struct {
	Sender    *Account `bson:"sender" json:"sender"`
	Recipient *Account `bson:"recipient" json:"recipient"`

	TimeOfLeaving       time.Time `bson:"timeofleaving" json:"timeofleaving"`
	TimeOfComing        time.Time `bson:"timeofcoming" json:"timeofcoming"`
	Amount              int64     `bson:"amount" json:"amount"`
	TransactionCurrency string    `bson:"transactioncurrency" json:"transactioncurrency"`

	TransactionFee int64 `bson:"transactionfee" json:"transactionfee"`

	TransactionHash string `bson:"transactionhash" json:"transactionhash"`

	Status TransactionStatus `bson:"transactionstatus" json:"transactionstatus"`

	Path []string `bson:"path" json:"path"`

	OnNetworkID uint32 `bson:"onNetworkID" json:"onNetworkID"`
}

// TODO
func NewTransaction(sender *Account, recipient *Account, amount int64, currency string, fee int64, onNetworkID uint32) *Transaction {
	transaction := new(Transaction)
	transaction.Sender = sender
	transaction.Recipient = recipient
	transaction.TimeOfLeaving = time.Now()
	transaction.Amount = amount
	transaction.TransactionCurrency = currency
	transaction.TransactionFee = fee
	transaction.OnNetworkID = onNetworkID
	return transaction
}

func (transaction *Transaction) SetPath(network *BankConnectionsGraph) {
	route, err := network.FindRoute(transaction.Sender.MaintingBankBIC, transaction.Recipient.MaintingBankBIC)
	if err != nil {
		log.Fatal(err.Error())
	}
	transaction.Path = route
}

// TODO
func (transaction *Transaction) Send() {

}

func (transaction *Transaction) Froze() {

}

func (transaction *Transaction) Unfroze() {

}

func (transaction *Transaction) Realise() {

}
