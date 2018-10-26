package login

type LoginAccountDrt struct {
	SessionToken *string `json:"token"`
	ErrorCode    *int    `json:"errorCode"`
	Error        bool    `json:"error"`
}

func newLoginAccountDrt(err bool, errorCode *int, sessionToken *string) *LoginAccountDrt {
	drt := new(LoginAccountDrt)
	drt.Error = err
	drt.ErrorCode = errorCode
	drt.SessionToken = sessionToken
	return drt
}
