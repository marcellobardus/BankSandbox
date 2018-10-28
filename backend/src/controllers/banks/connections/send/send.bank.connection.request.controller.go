package send

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func sendBankConnectionRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	var sendBankConnectionRequestDto SendBankConnectionRequestDto

	if err := json.NewDecoder(req.Body).Decode(&sendBankConnectionRequestDto); err != nil {
		code := 401
		res := NewSendBankConnectionRequestDrt(true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	// Identify sending bank

	sendingBank, err := database.DbConnection.GetBankByBIC(sendBankConnectionRequestDto.SenderBIC)

	if err != nil {
		code := 402
		res := NewSendBankConnectionRequestDrt(true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	// Authorize the sender

	privateKey := w.Header().Get("privateKey")
	if privateKey != sendingBank.PrivateKey {
		code := 403
		res := NewSendBankConnectionRequestDrt(true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if len(sendingBank.OwnersProfiles) != len(sendBankConnectionRequestDto.OTPs) {
		code := 404
		res := NewSendBankConnectionRequestDrt(true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	for i := 0; i < len(sendingBank.OwnersProfiles); i++ {
		auth := false
		for j := 0; j < len(sendBankConnectionRequestDto.OTPs); j++ {
			if authorized, _ := sendingBank.OwnersProfiles[i].Authenticate(sendBankConnectionRequestDto.OTPs[j]); authorized {
				auth = true
			}
		}
		if !auth {
			code := 405
			res := NewSendBankConnectionRequestDrt(true, &code)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	}

	// Indentify recipient bank

	recipientBank, err := database.DbConnection.GetBankByBIC(sendBankConnectionRequestDto.RecipientBIC)

	// create a connection request

	request := datamodels.NewBankConnectionRequest(sendingBank.BIC, recipientBank.BIC)

	// update banks requests

	sendingBank.OutcomingConnectionRequests = append(sendingBank.OutcomingConnectionRequests, request)
	recipientBank.IncomingConnectionRequests = append(recipientBank.IncomingConnectionRequests, request)

	// TODO update the banks connections requests in mongodb

	if err := database.DbConnection.UpdateBankByBIC(sendingBank.BIC, sendingBank); err != nil {
		code := 407
		res := NewSendBankConnectionRequestDrt(true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		log.Println(err)
		return
	}

	if err := database.DbConnection.UpdateBankByBIC(recipientBank.BIC, recipientBank); err != nil {
		code := 408
		res := NewSendBankConnectionRequestDrt(true, &code)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		log.Println(err)
		return
	}

	res := NewSendBankConnectionRequestDrt(false, nil)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return
}

// SendBankConnectionRequestController returns a bank send connection request controller
func SendBankConnectionRequestController() *utils.Controller {
	sendBankConnectionRequestController := utils.NewController("banks/connections/send-request", "POST", sendBankConnectionRequest)
	return sendBankConnectionRequestController
}
