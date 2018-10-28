package datamodels

import (
	"time"
)

type BankConnection struct {
	ConnectedBIC     string        `bson:"connectedBIC" json:"connectedBIC"`
	TransferTime     time.Duration `bson:"transferTime" json:"transferTime"`
	TransferTimeUnit interface{}   `bson:"transferTimeUnit" json:"transferTimeUnit"`
}

func NewBankConnection(connectedBIC string, transferTime time.Duration, transferTimeUnit interface{}) *BankConnection {
	bankConnection := new(BankConnection)
	bankConnection.ConnectedBIC = connectedBIC
	bankConnection.TransferTime = transferTime
	bankConnection.TransferTimeUnit = transferTimeUnit
	return bankConnection
}
