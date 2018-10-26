package create

import (
	"encoding/json"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func createAccount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	var createAccountDto CreateAccountDto

	if err := json.NewDecoder(req.Body).Decode(&createAccountDto); err != nil {
		code := 101
		res := newCreateAccountDrt(true, nil, nil, nil, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	account := datamodels.NewAccount(
		createAccountDto.Name,
		createAccountDto.Surname,
		createAccountDto.Mail,
		createAccountDto.PhoneNumber,
		createAccountDto.SocialInsuranceID,
		createAccountDto.PasswordHash)

	// TODO
	// account.setLoginID()
	account.SetOPT()
	account.Session = datamodels.NewSession(30)

	if err := database.DbConnection.InsertAccount(account); err != nil {
		code := 102
		res := newCreateAccountDrt(true, nil, nil, nil, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	sessionToken := &account.Session.Token
	otpSecret := &account.OTP.Secret
	loginID := &account.LoginID

	res := newCreateAccountDrt(false, sessionToken, otpSecret, loginID, nil)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return

}

// CreateAccountController returns an accounts creation controller
func CreateAccountController() *utils.Controller {
	createAccountController := utils.NewController("accounts/create", "POST", createAccount)
	return createAccountController
}
