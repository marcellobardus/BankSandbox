package controllers

import (
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/banks"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/accounts"

	"github.com/gorilla/mux"
)

func NewAppController() http.Handler {
	router := mux.NewRouter()

	// Account endpoints handlers

	router.HandleFunc(
		accounts.CreateAccountsController.Path,
		accounts.CreateAccountsController.HandlerFunc).Methods(accounts.CreateAccountsController.Method)
	router.HandleFunc(
		accounts.LoginAccountsController.Path,
		accounts.LoginAccountsController.HandlerFunc).Methods(accounts.LoginAccountsController.Method)

	// Banks endpoints handlers

	router.HandleFunc(
		banks.CreateBankController.Path,
		banks.CreateBankController.HandlerFunc).Methods(banks.CreateBankController.Method)
	router.HandleFunc(
		banks.SendBankConnectionRequestController.Path,
		banks.SendBankConnectionRequestController.HandlerFunc).Methods(banks.SendBankConnectionRequestController.Method)
	router.HandleFunc(
		banks.AnswerBankConnectionRequestController.Path,
		banks.AnswerBankConnectionRequestController.HandlerFunc).Methods(banks.AnswerBankConnectionRequestController.Method)
	return router
}
