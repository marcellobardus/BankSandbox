package datamodels

import (
	"errors"
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

	Path []*Bank `bson:"path" json:"path"`

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
func (transaction *Transaction) Send() error {
	var senderWalletID, recipientWalletID int
	// Check if recipient has a wallet which matches the transaction's currency
	ok := false
	for i := 0; i < len(transaction.Recipient.Wallets); i++ {
		if transaction.TransactionCurrency == transaction.Recipient.Wallets[i].Currency {
			ok = true
			recipientWalletID = i
		}
	}

	if !ok {
		errors.New("The transaction's recipient has not a wallet with the correct currency")
	}

	// Check if sender has a wallet which matches the transaction's currency
	ok = false
	for i := 0; i < len(transaction.Sender.Wallets); i++ {
		if transaction.TransactionCurrency == transaction.Sender.Wallets[i].Currency {
			ok = true
			senderWalletID = i
		}
	}

	if !ok {
		errors.New("The transaction's sender has not a wallet with the correct currency")
	}
	// Check if sender can send the transaction
	if transaction.Amount > transaction.Sender.Wallets[senderWalletID].Balance {
		return errors.New("Transaction's sender hasn't enough balance")
	}

	// Decrease sender's balance
	transaction.Sender.Wallets[senderWalletID].DecreaseBalance(transaction.Amount)

	// Increase recipient's incoming transfers balance
	transaction.Recipient.Wallets[recipientWalletID].IncreaseIncomingBalance(transaction.Amount)

	// Set time of coming

	// Check if path is defined

	if len(transaction.Path) == 0 || transaction.Path == nil {
		return errors.New("Transaction path is undefined")
	}

	// TODO set the time in base of the connections in the path

	var transactionDuration int64

	for i := 0; i < len(transaction.Path); i++ {
		for j := 0; i < len(transaction.Path[i].Connections); j++ {
			if transaction.Path[i].Connections[j].ConnectedBIC == transaction.Path[i].BIC {
				transactionDuration += int64(transaction.Path[i].Connections[j].TransferTime)
			}
		}
	}

	transaction.TimeOfComing = time.Now().Add(time.Duration(transactionDuration) * time.Minute)

	return nil
}

// TODO
func (transaction *Transaction) Froze() {

}

// TODO
func (transaction *Transaction) Unfroze() {

}

// TODO
func (transaction *Transaction) Realise() {

}
