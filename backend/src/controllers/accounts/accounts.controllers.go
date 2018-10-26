package accounts

import (
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/accounts/create"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/accounts/login"
)

var CreateAccountsController = create.CreateAccountController()
var LoginAccountsController = login.LoginAccountController()
