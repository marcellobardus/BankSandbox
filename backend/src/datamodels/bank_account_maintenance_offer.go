package datamodels

import (
	"errors"
)

type BankAccountMaintenanceOffer struct {
	SendingBankBIC                    string `bson:"sendingBankBIC" json:"sendingBankBIC"`
	RecipientAccountSocialInsuranceID string `bson:"recipientAccountSocialInsuranceID" json:"recipientAccountSocialInsuranceID"`
	MonthlyFee                        int64  `bson:"monthlyFee" json:"monthlyFee"`
	SingleTransferFee                 int64  `bson:"singleTransferFee" json:"singleTransferFee"`
	FeesCurrency                      string `bson:"feesCurrency" json:"feesCurrency"`
}

func NewBankAccountMaintenanceOffer(
	sendingBankBIC string,
	monthlyFee int64,
	singleTransferFee int64,
	feesCurrency string) *BankAccountMaintenanceOffer {
	offer := new(BankAccountMaintenanceOffer)
	offer.SendingBankBIC = sendingBankBIC
	offer.MonthlyFee = monthlyFee
	offer.SingleTransferFee = singleTransferFee
	offer.FeesCurrency = feesCurrency
	return offer
}

func (offer *BankAccountMaintenanceOffer) Send(recipientAccount *Account) error {
	for i := 0; i < len(recipientAccount.AccountMaintenanceOffers); i++ {
		if recipientAccount.AccountMaintenanceOffers[i].SendingBankBIC == offer.SendingBankBIC {
			return errors.New("Offer from this bank has been already sent to this user")
		}
	}
	offer.RecipientAccountSocialInsuranceID = recipientAccount.SocialInsuranceID
	recipientAccount.AccountMaintenanceOffers = append(recipientAccount.AccountMaintenanceOffers, offer)
	return nil
}

func (offer *BankAccountMaintenanceOffer) CancelOffer(recipientAccount *Account) error {
	for i := 0; i < len(recipientAccount.AccountMaintenanceOffers); i++ {
		if recipientAccount.AccountMaintenanceOffers[i].SendingBankBIC == offer.SendingBankBIC {
			arraySize := len(recipientAccount.AccountMaintenanceOffers)
			recipientAccount.AccountMaintenanceOffers[i] = recipientAccount.AccountMaintenanceOffers[arraySize-1]
			recipientAccount.AccountMaintenanceOffers[arraySize-1] = nil
			recipientAccount.AccountMaintenanceOffers = recipientAccount.AccountMaintenanceOffers[:arraySize-1]
			return nil
		}
	}
	return errors.New("offer not found")
}

// TODO
func (offer *BankAccountMaintenanceOffer) Accept(account *Account) error {
	return nil
}

// TODO
func (offer *BankAccountMaintenanceOffer) Refuse(account *Account) error {
	return nil
}
