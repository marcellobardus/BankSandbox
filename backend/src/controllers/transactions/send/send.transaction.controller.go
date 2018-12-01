package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func sendTransaction(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	var sendTransactionDto SendTransactionDto

	if err := json.NewDecoder(req.Body).Decode(&sendTransactionDto); err != nil {
		code := 601
		res := newSendTransactionDrt(nil, nil, nil, true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}
}

func SendTransactionController() *utils.Controller {
	sendTransactionController := utils.NewController("transactions/send", "POST", sendTransaction)
	return sendTransactionController
}
