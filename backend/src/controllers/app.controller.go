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
	return router
}
