package login

type LoginAccountDrt struct {
	SessionToken *string `json:"token"`
	ErrorCode    *int    `json:"errorCode"`
	Error        bool    `json:"error"`
}

func newLoginAccountDrt(err bool, errorCode *int, sessionToken *string) *LoginAccountDrt {
	drt := new(LoginAccountDrt)
	drt.Error = err
	drt.SessionToken = sessionToken
	if errorCode != nil {
		code := *errorCode + 1000
		drt.ErrorCode = &code
	} else {
		drt.ErrorCode = errorCode
	}
	return drt
}
