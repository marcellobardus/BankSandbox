package login

type LoginAccountDto struct {
	LoginID      uint32 `json:"loginID"`
	PasswordHash string `json:"passwordHash"`
	OTP          string `json:"OTP"`
}
