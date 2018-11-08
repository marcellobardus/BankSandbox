package answer

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

func answerBankConnectionRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer req.Body.Close()

	// Identify bank

	privateKey := req.Header.Get("privateKey")

	bank, err := database.DbConnection.GetBankByPrivateKey(privateKey)

	if err != nil {
		code := 501
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		log.Println(err)
		return
	}

	if bank.PrivateKey != privateKey {
		code := 502
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	var answerBankConnectionRequestDto AnswerBankConnectionRequestDto

	if err := json.NewDecoder(req.Body).Decode(&answerBankConnectionRequestDto); err != nil {
		code := 503
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if len(bank.OwnersProfiles) != len(answerBankConnectionRequestDto.OTPs) {
		code := 504
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
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
			res := newAnswerBankConnectionRequestDrt(true, &code, false)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	}

	isBankAuthorizedToAcceptRequest := false
	var connection *datamodels.BankConnection

	for i := 0; i < len(bank.IncomingConnectionRequests); i++ {
		if answerBankConnectionRequestDto.SenderBIC == bank.IncomingConnectionRequests[i].Sender {
			connection = datamodels.NewBankConnection(bank.BIC, bank.IncomingConnectionRequests[i].TransferTime, bank.IncomingConnectionRequests[i].TransferTimeUnit, int64(bank.IncomingConnectionRequests[i].TransferFee))
			isBankAuthorizedToAcceptRequest = true
			break
		}
	}

	if !isBankAuthorizedToAcceptRequest {
		code := 506
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	// Identify sender bank

	senderBank, err := database.DbConnection.GetBankByBIC(answerBankConnectionRequestDto.SenderBIC)

	if err != nil {
		code := 507
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}
	if answerBankConnectionRequestDto.Accept {
		err := bank.AcceptConnectionRequest(senderBank.BIC, uint32(connection.TransferTime), connection.TransferTimeUnit, connection.TransferFee)
		if err != nil {
			code := 508
			res := newAnswerBankConnectionRequestDrt(true, &code, false)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	} else {
		err := bank.RefuseConnectionRequest(senderBank.BIC)
		if err != nil {
			code := 508
			res := newAnswerBankConnectionRequestDrt(true, &code, false)
			resJSON, _ := json.Marshal(res)
			w.Write(resJSON)
			return
		}
	}

	err = senderBank.DeleteSentConnectionRequest(bank.BIC)

	if err != nil {
		code := 510
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	err = database.DbConnection.UpdateBankByBIC(bank.BIC, bank)

	if err != nil {
		code := 511
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if !answerBankConnectionRequestDto.Accept {
		res := newAnswerBankConnectionRequestDrt(false, nil, answerBankConnectionRequestDto.Accept)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	senderBank.Connections = append(senderBank.Connections, connection)

	err = database.DbConnection.UpdateBankByBIC(senderBank.BIC, senderBank)

	if err != nil {
		code := 512
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	graph, err := database.DbConnection.GetGraphByID(0)

	if err != nil {
		code := 513
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	routeDuration := senderBank.Connections[len(senderBank.Connections)-1].TransferTime
	routeFee := senderBank.Connections[len(senderBank.Connections)-1].TransferFee
	route := datamodels.NewBankConnectionRoute(senderBank.BIC, bank.BIC, uint8(routeFee), routeDuration)
	if err := graph.CreateNewRoute(route); err != nil {
		code := 514
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	if err := database.DbConnection.UpdateGraph(graph); err != nil {
		code := 515
		res := newAnswerBankConnectionRequestDrt(true, &code, false)
		resJSON, _ := json.Marshal(res)
		w.Write(resJSON)
		return
	}

	res := newAnswerBankConnectionRequestDrt(false, nil, answerBankConnectionRequestDto.Accept)
	resJSON, _ := json.Marshal(res)
	w.Write(resJSON)
	return
}

// AnswerBankConnectionRequestController returns an bank answer connection request controller
func AnswerBankConnectionRequestController() *utils.Controller {
	answerBankConnectionRequestController := utils.NewController("banks/connections/answer-request", "POST", answerBankConnectionRequest)
	return answerBankConnectionRequestController
}
