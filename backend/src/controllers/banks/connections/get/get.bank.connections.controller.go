package get

import (
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func getBankConnections(w http.ResponseWriter, req *http.Request) {

}

// GetBankConnectionsController returns a bank get connections controller
func GetBankConnectionsController() *utils.Controller {
	getBankConnectionsController := utils.NewController("banks/connections/get", "GET", getBankConnections)
	return getBankConnectionsController
}
