package banks

import (
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/banks/connections"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/banks/create"
)

var CreateBankController = create.CreateBankController()
var SendBankConnectionRequestController = connections.SendBankConnectionRequestController
var AnswerBankConnectionRequestController = connections.AnswerBankConnectionRequestController
