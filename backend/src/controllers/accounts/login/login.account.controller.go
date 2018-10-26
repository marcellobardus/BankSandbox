package login

import (
	"encoding/json"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func loginAccount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	var loginAccountDto LoginAccountDto

	if err := json.NewDecoder(req.Body).Decode(&loginAccountDto); err != nil {
		code := 201
		res := newLoginAccountDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	account, err := database.DbConnection.GetAccountByLoginID(loginAccountDto.LoginID)

	if err != nil {
		code := 202
		res := newLoginAccountDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if account.PasswordHash != loginAccountDto.PasswordHash {
		code := 203
		res := newLoginAccountDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	authenticated, err := account.OTP.Authenticate(loginAccountDto.OTP)

	if err != nil {
		code := 204
		res := newLoginAccountDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if !authenticated {
		code := 205
		res := newLoginAccountDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	account.Session = datamodels.NewSession(30)

	res := newLoginAccountDrt(false, nil, &account.Session.Token)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return

}

// LoginAccountController returns an accounts creation controller
func LoginAccountController() *utils.Controller {
	loginAccountController := utils.NewController("accounts/login", "POST", loginAccount)
	return loginAccountController
}
