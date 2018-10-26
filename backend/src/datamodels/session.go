package datamodels

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dchest/uniuri"
)

type Session struct {
	Token          string    `bson: "token" json:"token"`
	CreationTime   time.Time `bson:"creationTime" json:"creationTime"`
	ExpirationTime time.Time `bson:"expirationTime json:"expirationTime"`
}

func NewSession(expirationAfterMinutes uint16) *Session {
	// Initialize session
	session := new(Session)
	session.SetNewToken()
	session.CreationTime = time.Now()
	session.ExpirationTime = time.Now().Add(time.Minute * time.Duration(expirationAfterMinutes))
	return session
}

func (session *Session) IsSessionValid() bool {
	return session.ExpirationTime.Unix() <= time.Now().Unix()
}

func (session *Session) SetNewToken() {
	randomString := uniuri.New()
	sha256Hash := sha256.Sum256([]byte(randomString))
	session.Token = hex.EncodeToString(sha256Hash[:])
}
