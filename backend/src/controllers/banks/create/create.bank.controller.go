package create

import (
	"encoding/base32"
	"encoding/json"
	"net/http"

	"github.com/dgryski/dgoogauth"
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
		res := newCreateBankDrt(true, &code, nil, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	privateKey := utils.NewRandomSha512()

	var ownersProfiles = make([]*dgoogauth.OTPConfig, 0)

	for i := 0; i < int(createBankDto.OwnersNumber); i++ {
		ownerProfile := dgoogauth.OTPConfig{
			Secret:     base32.StdEncoding.EncodeToString([]byte(utils.NewRandomSha512())),
			WindowSize: 3, HotpCounter: 0}
		ownersProfiles = append(ownersProfiles, &ownerProfile)
	}

	bank := datamodels.NewBank(
		createBankDto.Name,
		createBankDto.CountryCode,
		createBankDto.BIC,
		privateKey,
		ownersProfiles,
		0)

	if err := database.DbConnection.InsertBank(bank); err != nil {
		code := 302 // Bank already exists
		res := newCreateBankDrt(true, &code, nil, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	var ownersProfilesSecrets = make([]string, 0)

	for i := 0; i < len(ownersProfiles); i++ {
		ownersProfilesSecrets = append(ownersProfilesSecrets, ownersProfiles[i].Secret)
	}

	graph, err := database.DbConnection.GetGraphByID(0)

	if err != nil {
		code := 303
		res := newCreateBankDrt(true, &code, nil, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	graph.PushNewBank(bank)

	if err := database.DbConnection.UpdateGraph(graph); err != nil {
		code := 304
		res := newCreateBankDrt(true, &code, nil, nil)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	message := "Because of security reasons the private key will not be delivired over the http/https protocol, please contact us"
	res := newCreateBankDrt(false, nil, &message, &ownersProfilesSecrets)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return

}

// CreateBankController returns an bank creation controller
func CreateBankController() *utils.Controller {
	createBankController := utils.NewController("banks/create", "POST", createBank)
	return createBankController
}
