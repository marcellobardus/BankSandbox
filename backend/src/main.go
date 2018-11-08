package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/controllers"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/daemons"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/database"
)

func main() {
	setDatabaseConnection()
	go daemons.TransactionsDaemon()
	startRest()
}

func setDatabaseConnection() {
	database.SetConnection("localhost", "banksandboxdb")
}

func startRest() {
	router := controllers.NewAppController()
	fmt.Println("Listening on port 8087")
	err := http.ListenAndServe(":8087", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
