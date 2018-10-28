package send

type SendBankConnectionRequestDrt struct {
	Error     bool `json:"error"`
	ErrorCode *int `json:"errorCode"`
}

func NewSendBankConnectionRequestDrt(err bool, errorCode *int) *SendBankConnectionRequestDrt {
	drt := new(SendBankConnectionRequestDrt)
	drt.Error = err
	if errorCode != nil {
		code := *errorCode + 1000
		drt.ErrorCode = &code
	} else {
		drt.ErrorCode = errorCode
	}
	return drt
}
