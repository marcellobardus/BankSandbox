package datamodels

import (
	"errors"
	"time"

	"github.com/dgryski/dgoogauth"
)

// Bank datamodel
type Bank struct {
	Name                        string                   `json:"name" bson:"name"`
	CountryCode                 string                   `json:"countryCode" bson:"countryCode"`
	BIC                         string                   `json:"bic" bson:"bic"`
	Connections                 []*BankConnection        `json:"connections" bson:"connections"`
	Customers                   []*Account               `json:"customers" bson:"customers"`
	PrivateKey                  string                   `json:"privateKey" bson:"privateKey"`
	OwnersProfiles              []*dgoogauth.OTPConfig   `json:"ownersProfiles" bson:"ownersProfiles"`
	IncomingConnectionRequests  []*BankConnectionRequest `json:"incomingConnectionRequests" bson:"incomingConnectionRequests"`
	OutcomingConnectionRequests []*BankConnectionRequest `json:"outcomingConnectionRequests" bson:"outcomingConnectionRequests"`
	BanksNetworkGraphID         uint32                   `bson:"banksNetworkGraphID" json:"banksNetworkGraphID"`
}

// NewBank creates a new bank
func NewBank(name string, countryCode string, bic string, privateKey string, ownersProfiles []*dgoogauth.OTPConfig, banksNetworkGraphID uint32) *Bank {
	bank := new(Bank)
	bank.CountryCode = countryCode
	bank.BIC = bic
	bank.Name = name
	bank.Connections = make([]*BankConnection, 0)
	bank.Customers = make([]*Account, 0)
	bank.OwnersProfiles = make([]*dgoogauth.OTPConfig, 0)
	bank.PrivateKey = privateKey
	bank.OwnersProfiles = ownersProfiles
	bank.BanksNetworkGraphID = banksNetworkGraphID
	return bank
}

func (sender *Bank) SendConnectionRequest(recipient *Bank, transferTime uint32, transferTimeUnit interface{}) {
	request := NewBankConnectionRequest(sender.BIC, recipient.BIC, transferTime, transferTimeUnit)
	sender.OutcomingConnectionRequests = append(sender.OutcomingConnectionRequests, request)
	recipient.IncomingConnectionRequests = append(recipient.IncomingConnectionRequests, request)
}

func (sender *Bank) DeleteSentConnectionRequest(recipient string) error {
	for i := 0; i < len(sender.OutcomingConnectionRequests); i++ {
		if sender.OutcomingConnectionRequests[i].Recipient == recipient {
			arraySize := len(sender.OutcomingConnectionRequests)
			sender.OutcomingConnectionRequests[i] = sender.OutcomingConnectionRequests[arraySize-1]
			sender.OutcomingConnectionRequests[arraySize-1] = nil
			sender.OutcomingConnectionRequests = sender.OutcomingConnectionRequests[:arraySize-1]
			return nil
		}
	}
	return errors.New("Wrong recipient argument")
}

func (recipient *Bank) DeleteDeliveredConnectionRequest(sender string) error {
	for i := 0; i < len(recipient.IncomingConnectionRequests); i++ {
		if recipient.IncomingConnectionRequests[i].Sender == sender {
			arraySize := len(recipient.IncomingConnectionRequests)
			recipient.IncomingConnectionRequests[i] = recipient.IncomingConnectionRequests[arraySize-1]
			recipient.IncomingConnectionRequests[arraySize-1] = nil
			recipient.IncomingConnectionRequests = recipient.IncomingConnectionRequests[:arraySize-1]
			return nil
		}
	}
	return errors.New("Wrong sender argument")
}

func (recipient *Bank) AcceptConnectionRequest(sender string, transferTime uint32, timeUnit interface{}, transferFee int64) error {
	for i := 0; i < len(recipient.IncomingConnectionRequests); i++ {
		if recipient.IncomingConnectionRequests[i].Sender == sender {
			recipient.IncomingConnectionRequests[i].Accept()
			newConnection := NewBankConnection(sender, time.Duration(transferTime), timeUnit, transferFee)
			recipient.Connections = append(recipient.Connections, newConnection)

			// Delete request

			arraySize := len(recipient.IncomingConnectionRequests)
			recipient.IncomingConnectionRequests[i] = recipient.IncomingConnectionRequests[arraySize-1]
			recipient.IncomingConnectionRequests[arraySize-1] = nil
			recipient.IncomingConnectionRequests = recipient.IncomingConnectionRequests[:arraySize-1]

			return nil
		}
	}

	return errors.New("Wrong from argument")
}

func (recipient *Bank) RefuseConnectionRequest(sender string) error {
	for i := 0; i < len(recipient.IncomingConnectionRequests); i++ {
		if recipient.IncomingConnectionRequests[i].Sender == sender {
			recipient.IncomingConnectionRequests[i].Refuse()

			// Delete request

			arraySize := len(recipient.IncomingConnectionRequests)
			recipient.IncomingConnectionRequests[i] = recipient.IncomingConnectionRequests[arraySize-1]
			recipient.IncomingConnectionRequests[arraySize-1] = nil
			recipient.IncomingConnectionRequests = recipient.IncomingConnectionRequests[:arraySize-1]

			return nil
		}
	}

	return errors.New("Wrong from argument")
}

func (bank *Bank) SendTransaction(
	senderAccount *Account,
	recipientAccount *Account,
	recipientBank *Bank,
	amount int64,
	fee int64,
	transactionCurrency string,
	network *BankConnectionsGraph) (*Transaction, error) {

	// Check if bank can send a transaction for the sender

	if bank.BIC != senderAccount.MaintingBankBIC {
		return nil, errors.New("Sending bank is not authorized to send a transaction for the sender")
	}

	// Check if recipientAccountsBank equals the recipientBank

	if recipientBank.BIC != recipientAccount.MaintingBankBIC {
		return nil, errors.New("recipientAccountsBank is not equal to recipientBank")
	}

	// Find a path to the recipientsBank

	// Check if the given networkID is equal to the banks network ID

	if bank.BanksNetworkGraphID != network.ID {
		return nil, errors.New("bank is not authorized to use this network")
	}

	// Create transaction and find path

	transaction := NewTransaction(senderAccount, recipientAccount, amount, transactionCurrency, fee, bank.BanksNetworkGraphID)

	// TODO
	transaction.SetPath(network)
	transaction.Send()

	return nil, nil
}
