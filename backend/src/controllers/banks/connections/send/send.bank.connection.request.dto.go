package send

type SendBankConnectionRequestDto struct {
	RecipientBIC string   `json:"recipientBIC"`
	SenderBIC    string   `json:"senderBIC"`
	OTPs         []string `json:"otps"`
}
