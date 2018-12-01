package transactions

// Header params
//@privateKey: string
//@session: string

type SendTransactionDto struct {
	Value            int64  `json:"value"`
	RecipientIBAN    string `json:"recipientIBAN"`
	RecipientBankBIC string `json:"recipientBankBIC"`
}
