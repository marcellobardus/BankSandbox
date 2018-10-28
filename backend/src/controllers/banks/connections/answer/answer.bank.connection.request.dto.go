package answer

type AnswerBankConnectionRequestDto struct {
	SenderBIC string   `json:"senderBIC"`
	Accept    bool     `json:"accept"`
	OTPs      []string `json:"otps"`
}
