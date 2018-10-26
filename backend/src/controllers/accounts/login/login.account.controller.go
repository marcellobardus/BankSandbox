package login

import (
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func loginAccount(w http.ResponseWriter, req *http.Request) {

}

// LoginAccountController returns an accounts creation controller
func LoginAccountController() *utils.Controller {
	createAccountController := utils.NewController("accounts/login", "POST", loginAccount)
	return createAccountController
}
