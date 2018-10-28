package datamodels

import (
	"errors"

	"github.com/dgryski/dgoogauth"
)

// Bank datamodel
type Bank struct {
	Name                        string                   `json:"name" bson:"name"`
	CountryCode                 string                   `json:"countryCode" bson:"countryCode"`
	BIC                         string                   `json:"bic" bson:"bic"`
	Connections                 []string                 `json:"connections" bson:"connections"`
	Customers                   []*Account               `json:"customers" bson:"customers"`
	PrivateKey                  string                   `json:"privateKey" bson:"privateKey"`
	OwnersProfiles              []*dgoogauth.OTPConfig   `json:"ownersProfiles" bson:"ownersProfiles"`
	IncomingConnectionRequests  []*BankConnectionRequest `json:"incomingConnectionRequests" bson:"incomingConnectionRequests"`
	OutcomingConnectionRequests []*BankConnectionRequest `json:"outcomingConnectionRequests" bson:"outcomingConnectionRequests"`
}

// NewBank creates a new bank
func NewBank(name string, countryCode string, bic string, privateKey string, ownersProfiles []*dgoogauth.OTPConfig) *Bank {
	bank := new(Bank)
	bank.CountryCode = countryCode
	bank.BIC = bic
	bank.Connections = make([]string, 0)
	bank.Customers = make([]*Account, 0)
	bank.OwnersProfiles = make([]*dgoogauth.OTPConfig, 0)
	bank.PrivateKey = privateKey
	bank.OwnersProfiles = ownersProfiles
	return bank
}

func (sender *Bank) SendConnectionRequest(recipient *Bank) {
	request := NewBankConnectionRequest(sender.BIC, recipient.BIC)
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

func (recipient *Bank) AcceptConnectionRequest(sender string) error {
	for i := 0; i < len(recipient.IncomingConnectionRequests); i++ {
		if recipient.IncomingConnectionRequests[i].Sender == sender {
			recipient.IncomingConnectionRequests[i].Accept()
			recipient.Connections = append(recipient.Connections, sender)

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
