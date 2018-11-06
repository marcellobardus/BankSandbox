package datamodels

import (
	"time"
)

type BankConnection struct {
	ConnectedBIC     string        `bson:"connectedBIC" json:"connectedBIC"`
	TransferTime     time.Duration `bson:"transferTime" json:"transferTime"`
	TransferTimeUnit interface{}   `bson:"transferTimeUnit" json:"transferTimeUnit"`
	TransferFee      int64         `bson:"transferFee" json:"transferFee"`
}

func NewBankConnection(connectedBIC string, transferTime time.Duration, transferTimeUnit interface{}, transferFee int64) *BankConnection {
	bankConnection := new(BankConnection)
	bankConnection.ConnectedBIC = connectedBIC
	bankConnection.TransferTime = transferTime
	bankConnection.TransferTimeUnit = transferTimeUnit
	bankConnection.TransferFee = transferFee
	return bankConnection
}
