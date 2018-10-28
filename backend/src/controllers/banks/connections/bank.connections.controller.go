package connections

import (
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/banks/connections/answer"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/banks/connections/send"
)

var SendBankConnectionRequestController = send.SendBankConnectionRequestController()
var AnswerBankConnectionRequestController = answer.AnswerBankConnectionRequestController()
