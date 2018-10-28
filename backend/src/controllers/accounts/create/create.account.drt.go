package create

type CreateAccountDrt struct {
	Error        bool    `json:"error"`
	SessionToken *string `json:"token"`
	OTPSecret    *string `json:"OTPSecret"`
	LoginID      *uint32 `json:"loginID"`
	ErrorCode    *int    `json:"errorCode"`
}

func newCreateAccountDrt(err bool,
	sessionToken *string,
	otpSecret *string,
	loginID *uint32,
	errorCode *int) *CreateAccountDrt {
	drt := new(CreateAccountDrt)
	drt.Error = err
	drt.SessionToken = sessionToken
	drt.OTPSecret = otpSecret
	drt.LoginID = loginID
	if errorCode != nil {
		code := *errorCode + 1000
		drt.ErrorCode = &code
	} else {
		drt.ErrorCode = errorCode
	}
	return drt
}
