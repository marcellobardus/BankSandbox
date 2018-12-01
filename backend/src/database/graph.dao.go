package database

import (
	"errors"
	"log"

	"github.com/spaghettiCoderIT/BankSandbox/backend/src/datamodels"
	"gopkg.in/mgo.v2/bson"
)

func (connection *Connection) InsertConnectionsGraph(graph *datamodels.BankConnectionsGraph) error {
	existingGraph, selectionErr := connection.GetGraphByID(graph.ID)
	if selectionErr != nil && selectionErr.Error() != "not found" {
		log.Fatal(selectionErr.Error())
	}

	if existingGraph != nil {
		return errors.New("Graph with the given ID already exists")
	}
	err := database.C(GraphsCollection).Insert(graph)
	return err
}

func (connection *Connection) GetAllGraphs() ([]*datamodels.BankConnectionsGraph, error) {
	var graphs []*datamodels.BankConnectionsGraph
	err := database.C(GraphsCollection).Find(bson.M{}).All(&graphs)
	return graphs, err
}

func (connection *Connection) GetGraphByID(id uint32) (*datamodels.BankConnectionsGraph, error) {
	var graph *datamodels.BankConnectionsGraph
	err := database.C(GraphsCollection).Find(bson.M{"id": id}).One(&graph)
	return graph, err
}

func (connection *Connection) UpdateGraph(graph *datamodels.BankConnectionsGraph) error {
	err := database.C(GraphsCollection).Update(bson.M{"id": graph.ID}, graph)
	return err
}
