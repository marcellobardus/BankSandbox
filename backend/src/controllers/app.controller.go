package controllers

import (
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers/accounts"

	"github.com/gorilla/mux"
)

func NewAppController() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(
		accounts.CreateAccountsController.Path,
		accounts.CreateAccountsController.HandlerFunc).Methods(accounts.CreateAccountsController.Method)
	return router
}
