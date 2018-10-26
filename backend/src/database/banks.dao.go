package database

import (
	"errors"
	"log"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"gopkg.in/mgo.v2/bson"
)

func (connection *Connection) InsertBank(bank *datamodels.Bank) error {
	existingBank, selectionError := connection.GetBankByBIC(bank.BIC)

	if selectionError != nil {
		log.Fatal(selectionError.Error())
	}

	if existingBank != nil {
		return errors.New("Bank already exists")
	}

	err := database.C(BanksCollection).Insert(bank)
	return err
}

func (connection *Connection) GetBankByCountryCodeAndBIC(countryCode string, bic string) (*datamodels.Bank, error) {
	var bank *datamodels.Bank
	err := database.C(BanksCollection).Find(bson.M{"countryCode": countryCode, "bic": bic}).One(&bank)
	return bank, err
}

func (connection *Connection) GetBankByBIC(bic string) (*datamodels.Bank, error) {
	var bank *datamodels.Bank
	err := database.C(BanksCollection).Find(bson.M{"bic": bic}).One(&bank)
	return bank, err
}
