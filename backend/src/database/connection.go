package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Connection defines the location of the database
type Connection struct {
	Server   string
	Database string
}

const (
	// BanksCollection cointains the name of banks collection name in mongoDB
	BanksCollection = "banks"
	// AccountsCollection cointains the name of accounts collection name in mongoDB
	AccountsCollection = "accounts"
	// RealisedTransactionsCollection cointains the name of realised_transactions collection name in mongoDB
	RealisedTransactionsCollection = "realised_transactions"
	// UnrealisedTransactionsCollection cointains the name of unrealised_transactions collection name in mongoDB
	UnrealisedTransactionsCollection = "unrealised_transactions"
	// GraphsCollection cointains the name of graphs collection name in mongoDB
	GraphsCollection = "graphs"
)

// Database allows to interact with the mongo collections
var database *mgo.Database
var DbConnection *Connection

// Connect establihes a new connection with mongodb
func (connection *Connection) Connect() {
	session, err := mgo.Dial(connection.Server)
	if err != nil {
		log.Fatal("dupa", err)
	}

	database = session.DB(connection.Database)
}

func SetConnection(server string, database string) {
	DbConnection = new(Connection)
	DbConnection.Database = database
	DbConnection.Server = server
	DbConnection.Connect()
}
