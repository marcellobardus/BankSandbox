package datamodels

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

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

	Wallets map[string]*Wallet
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
