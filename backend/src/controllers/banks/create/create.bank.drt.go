package create

type CreateBankDtr struct {
	Error            bool      `json:"error"`
	ErrorCode        *int      `json:"errorCode"`
	Message          *string   `json:"message"`
	OwnersOTPSecrets *[]string `json:"ownersOTPSecrets"`
}

func newCreateBankDrt(err bool, errorCode *int, message *string, ownersOTPSecrets *[]string) *CreateBankDtr {
	drt := new(CreateBankDtr)
	drt.Error = err
	drt.Message = message
	drt.OwnersOTPSecrets = ownersOTPSecrets
	if errorCode != nil {
		code := *errorCode + 1000
		drt.ErrorCode = &code
	} else {
		drt.ErrorCode = errorCode
	}
	return drt
}
