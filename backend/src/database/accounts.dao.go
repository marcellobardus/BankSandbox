package database

import (
	"errors"
	"log"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"gopkg.in/mgo.v2/bson"
)

func (connection *Connection) InsertAccount(account *datamodels.Account) error {
	existingAccount, selectionError := connection.GetAccountBySocialInsuranceID(account.SocialInsuranceID)

	if selectionError != nil && selectionError.Error() != "not found" {
		log.Fatal(selectionError.Error())
	}

	if existingAccount != nil {
		return errors.New("Account already exists")
	}

	err := database.C(AccountsCollection).Insert(account)
	return err
}

func (connection *Connection) GetAccountBySocialInsuranceID(id string) (*datamodels.Account, error) {
	var account *datamodels.Account
	err := database.C(AccountsCollection).Find(bson.M{"socialInsuranceID": id}).One(&account)
	return account, err
}

func (connection *Connection) GetAccountBySessionToken(sessionToken string) (*datamodels.Account, error) {
	var account *datamodels.Account
	err := database.C(AccountsCollection).Find(bson.M{"session.token": sessionToken}).One(&account)
	return account, err
}

func (connection *Connection) GetAccountByLoginID(loginID uint32) (*datamodels.Account, error) {
	var account *datamodels.Account
	err := database.C(AccountsCollection).Find(bson.M{"loginID": loginID}).One(&account)
	return account, err
}
