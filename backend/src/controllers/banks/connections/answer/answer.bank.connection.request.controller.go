package answer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func answerBankConnectionRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	// Identify bank

	privateKey := w.Header().Get("privateKey")

	bank, err := database.DbConnection.GetBankByPrivateKey(privateKey)

	if err != nil {
		code := 501
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		log.Println(err)
		return
	}

	if bank.PrivateKey != privateKey {
		code := 502
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	var answerBankConnectionRequestDto AnswerBankConnectionRequestDto

	if err := json.NewDecoder(req.Body).Decode(&answerBankConnectionRequestDto); err != nil {
		code := 503
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if len(bank.OwnersProfiles) != len(answerBankConnectionRequestDto.OTPs) {
		code := 504
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	for i := 0; i < len(bank.OwnersProfiles); i++ {
		auth := false
		for j := 0; j < len(answerBankConnectionRequestDto.OTPs); j++ {
			if authorized, _ := bank.OwnersProfiles[i].Authenticate(answerBankConnectionRequestDto.OTPs[j]); authorized {
				auth = true
			}
		}
		if !auth {
			code := 505
			res := NewAnswerBankConnectionRequestDrt(true, &code, false)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	}

	isBankAuthorizedToAcceptRequest := false

	for i := 0; i < len(bank.IncomingConnectionRequests); i++ {
		if answerBankConnectionRequestDto.SenderBIC == bank.IncomingConnectionRequests[i].Sender {
			isBankAuthorizedToAcceptRequest = true
			break
		}
	}

	if !isBankAuthorizedToAcceptRequest {
		code := 506
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	// Identify sender bank

	senderBank, err := database.DbConnection.GetBankByBIC(answerBankConnectionRequestDto.SenderBIC)

	if err != nil {
		code := 507
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}
	if answerBankConnectionRequestDto.Accept {
		err := bank.AcceptConnectionRequest(senderBank.BIC)
		if err != nil {
			code := 508
			res := NewAnswerBankConnectionRequestDrt(true, &code, false)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	} else {
		err := bank.RefuseConnectionRequest(senderBank.BIC)
		if err != nil {
			code := 508
			res := NewAnswerBankConnectionRequestDrt(true, &code, false)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	}

	err = bank.DeleteDeliveredConnectionRequest(senderBank.BIC)

	if err != nil {
		code := 509
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	err = senderBank.DeleteSentConnectionRequest(bank.BIC)

	if err != nil {
		code := 510
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	err = database.DbConnection.UpdateBankByBIC(bank.BIC, bank)

	if err != nil {
		code := 511
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	err = database.DbConnection.UpdateBankByBIC(senderBank.BIC, senderBank)

	if err != nil {
		code := 512
		res := NewAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	res := NewAnswerBankConnectionRequestDrt(false, nil, answerBankConnectionRequestDto.Accept)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return
}

// AnswerBankConnectionRequestController returns an bank answer connection request controller
func AnswerBankConnectionRequestController() *utils.Controller {
	answerBankConnectionRequestController := utils.NewController("banks/connections/answer-request", "POST", answerBankConnectionRequest)
	return answerBankConnectionRequestController
}
