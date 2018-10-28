package send

type SendBankConnectionRequestDto struct {
	RecipientBIC string   `json:"recipientBic"`
	SenderBIC    string   `json:"senderBIC"`
	OTPs         []string `json:"otps"`
}
