package datamodels

import "github.com/dgryski/dgoogauth"

// Bank datamodel
type Bank struct {
	Name           string                 `json:"name" bson:"name"`
	CountryCode    string                 `json:"countryCode" bson:"countryCode"`
	BIC            string                 `json:"bic" bson:"bic"`
	Connections    []uint32               `json:"connections" bson:"connections"`
	Customers      []*Account             `json:"customers" bson:"customers"`
	PrivateKey     string                 `json:"privateKey" bson:"privateKey"`
	OwnersProfiles []*dgoogauth.OTPConfig `json:"ownersProfiles" bson:"ownersProfiles"`
}

// NewBank creates a new bank
func NewBank(name string, countryCode string, bic string, privateKey string) *Bank {
	bank := new(Bank)
	bank.CountryCode = countryCode
	bank.BIC = bic
	bank.Connections = make([]uint32, 0)
	bank.Customers = make([]*Account, 0)
	bank.OwnersProfiles = make([]*dgoogauth.OTPConfig, 0)
	bank.PrivateKey = privateKey
	return bank
}
