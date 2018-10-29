package datamodels

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hash/adler32"
	"math/rand"
	"strconv"
	"time"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"

	"github.com/dgryski/dgoogauth"
)

// Account is a representation of a person who own a wallet and it's personal data
type Account struct {
	Name              string `bson:"name" json:"name"`
	Surname           string `bson:"surname" json:"surname"`
	SocialInsuranceID string `bson:"socialInsuranceID" json:"socialInsuranceID"`
	Mail              string `bson:"mail" json:"mail"`
	PhoneNumber       string `bson:"phoneNumber" json:"phoneNumber"`

	LoginID      uint32               `bson:"loginID" json:"loginID"`
	PasswordHash string               `bson:"passwordHash" json:"passwordHash"`
	OTP          *dgoogauth.OTPConfig `bson:"otp" json:"otp"`
	Session      *Session             `bson:"session" json:"session"`

	RegistrationDate time.Time `bson:"registrationDate" json:"registrationDate"`

	Wallets []*Wallet `bson:"wallets" json:"wallets"`

	AccountMaintenanceOffers []*BankAccountMaintenanceOffer `bson:"accountMaintenanceOffers" json:"accountMaintenanceOffers"`

	MaintingBankBIC string `bson:"maintingBankBIC" json:"maintingBankBIC"`
}

// NewAccount return a new account
func NewAccount(name string, surname string, mail string, phonenumber string, socialInsuranceID string, passwordHash string) *Account {
	account := new(Account)
	account.Name = name
	account.Surname = surname
	account.SocialInsuranceID = socialInsuranceID
	account.Mail = mail
	account.PhoneNumber = phonenumber

	account.LoginID = 0
	account.Session = NewSession(30)
	account.PasswordHash = passwordHash

	account.Wallets = make([]*Wallet, 0)

	account.AccountMaintenanceOffers = make([]*BankAccountMaintenanceOffer, 0)

	return account
}

func (account *Account) SetOPT() {
	rand.NewSource(time.Now().UnixNano())
	randomInt := rand.Intn(999999999999-9999) + 9999
	md5Hash := md5.Sum([]byte(strconv.Itoa(randomInt)))
	md5HashToString := hex.EncodeToString(md5Hash[:])
	account.OTP = &dgoogauth.OTPConfig{
		Secret:      md5HashToString,
		WindowSize:  3,
		HotpCounter: 0,
	}
}

func (account *Account) SetLoginID() error {
	if account.LoginID != 0 {
		return errors.New("Account's LoginID is already set")
	}

	randomSha521 := utils.NewRandomSha512()
	concatenatedString := utils.ConcatenateStrings(randomSha521, account.MaintingBankBIC, account.SocialInsuranceID)
	md5Hash := md5.Sum([]byte(concatenatedString))
	md5HashToString := hex.EncodeToString(md5Hash[:])
	adler32Hash := adler32.Checksum([]byte(md5HashToString))
	account.LoginID = adler32Hash

	return nil
}

func (account *Account) AppendWallet(wallet *Wallet) error {
	for i := 0; i < len(account.Wallets); i++ {
		if account.Wallets[i].Currency == wallet.Currency {
			return errors.New("Wallet with the given currency already exists")
		}
	}
	account.Wallets = append(account.Wallets, wallet)
	return nil
}

func (account *Account) DeleteWallet(currency string) error {
	for i := 0; i < len(account.Wallets); i++ {
		if account.Wallets[i].Currency == currency {
			arraySize := len(account.Wallets)
			account.Wallets[i] = account.Wallets[arraySize-1]
			account.Wallets[arraySize-1] = nil
			account.Wallets = account.Wallets[:arraySize-1]
			return nil
		}
	}
	return errors.New("Wallet with the given could not be deleted because it does not exist")
}

// TODO
func (account *Account) SetMaintingBank(bankBic string) error {
	return nil
}
