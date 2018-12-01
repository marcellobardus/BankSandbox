package datamodels

import (
	"errors"
	"log"
	"time"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

type TransactionStatus string

const (
	// Unconfirmed transaction cannot be send bcause its fee is undefined
	Unconfirmed TransactionStatus = "unconfirmed"
	// Sent => Transction is waiting for the nearest session
	Sent TransactionStatus = "sent"
	// Realised => Sender and Recipient wallets balances are updated
	Realised TransactionStatus = "realised"
	// Cancelled => Transaction is cancelled
	Cancelled TransactionStatus = "cancelled"
	// Frozen => Transaction is frozen
	Frozen TransactionStatus = "frozen"
)

type Transaction struct {
	Sender    *Account `bson:"sender" json:"sender"`
	Recipient *Account `bson:"recipient" json:"recipient"`

	SenderIBAN    string `bson:"senderIBAN" json:"senderIBAN"`
	RecipientIBAN string `bson:"recipientIBAN" json:"recipientIBAN"`

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

func (transaction *Transaction) SetIBANs() error {
	// Set the sender's IBAN
	for i := 0; i < len(transaction.Sender.Wallets); i++ {
		if transaction.Sender.Wallets[i].Currency == transaction.TransactionCurrency {
			transaction.SenderIBAN = transaction.Sender.Wallets[i].IBAN
		}
	}
	if transaction.SenderIBAN == "" {
		return errors.New("Could not find the sender's wallet which corresponds to the transaction currency")
	}

	// Set recipient's IBAN
	for i := 0; i < len(transaction.Recipient.Wallets); i++ {
		if transaction.Recipient.Wallets[i].Currency == transaction.TransactionCurrency {
			transaction.RecipientIBAN = transaction.Recipient.Wallets[i].IBAN
		}
	}
	if transaction.RecipientIBAN == "" {
		return errors.New("Could not find the recipient's wallet which corresponds to the transaction currency")
	}
	return nil
}

func (transaction *Transaction) SetPath(network *BankConnectionsGraph) {
	route, err := network.FindRoute(transaction.Sender.MaintingBankBIC, transaction.Recipient.MaintingBankBIC)
	if err != nil {
		log.Fatal(err.Error())
	}
	transaction.Path = route
}

func (transaction *Transaction) SetHash() error {
	randomSha256 := utils.NewRandomSha256()
	// check if sender and reicpient IBANs are set

	if transaction.SenderIBAN == "" {
		return errors.New("Transaction sender's IBAN has not been set")
	}

	if transaction.RecipientIBAN == "" {
		return errors.New("Transaction recipient's IBAN has not been set")
	}

	meta := utils.ConcatenateStrings(
		string(transaction.TimeOfLeaving.Unix()),
		string(transaction.Amount),
		transaction.SenderIBAN,
		transaction.RecipientIBAN,
		randomSha256)

	transaction.TransactionHash = meta

	return nil
}

// TODO setting time of coming
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
		return errors.New("The transaction's recipient has not a wallet with the correct currency")
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
		return errors.New("The transaction's sender has not a wallet with the correct currency")
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

func (transaction *Transaction) Realise() error {
	// When a time struct.Unix() returns a value < 0 it means that the date is still undefined
	if transaction.TimeOfComing.Unix() < 0 {
		return errors.New("Cannot realise an unsent transaction")
	}

	// Check if transaction can be realised
	if time.Now().Unix() < transaction.TimeOfComing.Unix() {
		return errors.New("It's to early to realise this transaction")
	}

	// Check if recipient has a wallet which matches the transaction's currency
	var recipientWalletID int
	ok := false
	for i := 0; i < len(transaction.Recipient.Wallets); i++ {
		if transaction.TransactionCurrency == transaction.Recipient.Wallets[i].Currency {
			ok = true
			recipientWalletID = i
		}
	}
	if !ok {
		return errors.New("The transaction's recipient has not a wallet with the correct currency")
	}

	transaction.Recipient.Wallets[recipientWalletID].DecreaseIncomingBalance(transaction.Amount)
	transaction.Recipient.Wallets[recipientWalletID].IncreaseBalance(transaction.Amount)

	transaction.SetHash()
	transaction.Status = Realised

	return nil
}
