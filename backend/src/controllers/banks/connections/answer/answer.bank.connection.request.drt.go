package answer

type AnswerBankConnectionRequestDrt struct {
	Error     bool `json:"error"`
	ErrorCode *int `json:"errorCode"`
	Accepted  bool `json:"accepted`
}

func NewAnswerBankConnectionRequestDrt(err bool, errorCode *int, accepted bool) *AnswerBankConnectionRequestDrt {
	drt := new(AnswerBankConnectionRequestDrt)
	drt.Error = err
	drt.Accepted = accepted
	if errorCode != nil {
		code := *errorCode + 1000
		drt.ErrorCode = &code
	} else {
		drt.ErrorCode = errorCode
	}
	return drt
}
