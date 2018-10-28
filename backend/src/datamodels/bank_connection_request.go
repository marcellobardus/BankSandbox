package datamodels

import (
	"time"
)

type BankConnectionRequestStatus string

const (
	Pending  BankConnectionRequestStatus = "pending"
	Accepted BankConnectionRequestStatus = "accepted"
	Refused  BankConnectionRequestStatus = "refused"
)

type BankConnectionRequest struct {
	// From property is specified by BIC
	Sender string `json:"from" bson:"from"`
	// To property is specified by BIC
	Recipient string `json:"to" bson:"to"`

	OutcomeDate     time.Time `json:"outcomeDate" bson:"outcomeDate"`
	AcceptationDate time.Time `json:"acceptationDate" bson:"acceptationDate"`

	Status BankConnectionRequestStatus `json:"status" bson:"status"`
}

func NewBankConnectionRequest(sender string, recipient string) *BankConnectionRequest {
	request := new(BankConnectionRequest)
	request.Sender = sender
	request.Recipient = recipient
	request.OutcomeDate = time.Now()
	request.Status = Pending
	return request
}

func (bankConnectionRequest *BankConnectionRequest) Accept() {
	bankConnectionRequest.Status = Accepted
	bankConnectionRequest.AcceptationDate = time.Now()
}

func (bankConnectionRequest *BankConnectionRequest) Refuse() {
	bankConnectionRequest.Status = Refused
	bankConnectionRequest.AcceptationDate = time.Now()
}
