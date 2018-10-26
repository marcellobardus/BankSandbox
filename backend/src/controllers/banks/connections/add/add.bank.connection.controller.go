package add

import (
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func addBankConnection(w http.ResponseWriter, req *http.Request) {

}

// AddBankConnectionController returns a bank add connection controller
func AddBankConnectionController() *utils.Controller {
	addBankConnectionController := utils.NewController("banks/connections/add", "POST", addBankConnection)
	return addBankConnectionController
}
