package datamodels

// Bank datamodel
type Bank struct {
	ID          uint32     `json:"id" bson:"id"`
	Name        string     `json:"name" bson:"name"`
	CountryCode string     `json:"countryCode" bson:"countryCode"`
	BIC         uint16     `json:"bic" bson:"bic"`
	Connections []uint32   `json:"connections" bson:"connections"`
	Customers   []*Account `json:"customers" bson:"customers"`
}

// NewBank creates a new bank
func NewBank(id uint32, name string, countryCode string, bic uint16) *Bank {
	bank := new(Bank)
	bank.ID = id
	bank.CountryCode = countryCode
	bank.BIC = bic
	bank.Connections = make([]uint32, 0)
	bank.Customers = make([]*Account, 0)
	return bank
}
