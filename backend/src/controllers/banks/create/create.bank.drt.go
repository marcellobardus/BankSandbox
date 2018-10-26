package create

type CreateBankDtr struct {
	Error     bool    `json:"error"`
	ErrorCode *int    `json:"errorCode"`
	Message   *string `json:"message"`
}

func newCreateBankDrt(err bool, errorCode *int, message *string) *CreateBankDtr {
	drt := new(CreateBankDtr)
	drt.Error = err
	drt.ErrorCode = errorCode
	drt.Message = message
	return drt
}
