package create

import (
	"encoding/json"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func createBank(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	var createBankDto CreateBankDto

	if err := json.NewDecoder(req.Body).Decode(&createBankDto); err != nil {
		code := 301
		res := newCreateBankDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	privateKey := utils.NewRandomSha512()

	bank := datamodels.NewBank(createBankDto.Name, createBankDto.CountryCode, createBankDto.BIC, privateKey)

	if err := database.DbConnection.InsertBank(bank); err != nil {
		code := 302
		res := newCreateBankDrt(true, &code, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	message := "Because of security reasons the private key will not be delivired over the http/https protocol, please contact us"
	res := newCreateBankDrt(false, nil, &message)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return

}

// CreateBankController returns an bank creation controller
func CreateBankController() *utils.Controller {
	createBankController := utils.NewController("banks/create", "POST", createBank)
	return createBankController
}
