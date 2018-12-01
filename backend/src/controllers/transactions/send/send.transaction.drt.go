package transactions

type SendTransactionDrt struct {
	Success   *bool   `json:"success"`
	TxHash    *string `json:"txHash"`
	Fees      *uint64 `json:"fees"`
	Error     bool    `json:"error"`
	ErrorCode *int    `json:"errorCode"`
}

func newSendTransactionDrt(success *bool, txHash *string, fees *uint64, err bool, errCode *int) *SendTransactionDrt {
	res := new(SendTransactionDrt)
	res.Success = success
	res.TxHash = txHash
	res.Fees = fees
	res.Error = err
	res.ErrorCode = errCode
	return res
}
