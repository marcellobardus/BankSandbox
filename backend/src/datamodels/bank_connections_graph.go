package datamodels

import (
	"errors"
	"time"

	"github.com/RyanCarrier/dijkstra"
	"github.com/spaghettiCoderIT/BankSandbox/backend/src/utils"
)

type BankConnectionRoute struct {
	FromBIC  string        `bson:"from" json:"from"`
	ToBIC    string        `bson:"to" json:"to"`
	FeesCost uint8         `bson:"feesCost" json:"feesCost"`
	TimeCost time.Duration `bson:"timeCost" json:"timeCost"`
}

func NewBankConnectionRoute(fromBIC string, toBIC string, feesCost uint8, timeCost time.Duration) *BankConnectionRoute {
	route := new(BankConnectionRoute)
	route.FromBIC = fromBIC
	route.ToBIC = toBIC
	route.FeesCost = feesCost
	route.TimeCost = timeCost
	return route
}

type dijkstraBy string

const (
	Time     dijkstraBy = "time"
	Fees     dijkstraBy = "fees"
	Distance dijkstraBy = "distance"
)

type BankConnectionsGraph struct {
	ID     uint32                 `bson:""id json:"id"`
	Banks  map[string]*Bank       `bson:"banks" json:"banks"`
	Routes []*BankConnectionRoute `bson:"routes" json:"routes"`
}

func NewBankConnectionsGraph(id uint32) *BankConnectionsGraph {
	graph := new(BankConnectionsGraph)
	graph.Banks = make(map[string]*Bank, 0)
	graph.Routes = make([]*BankConnectionRoute, 0)
	graph.ID = id
	return graph
}

func (graph *BankConnectionsGraph) PushNewBank(bank *Bank) {
	graph.Banks[bank.BIC] = bank
}

func (graph *BankConnectionsGraph) CreateNewRoute(route *BankConnectionRoute) error {
	banksBICs := mapKeyToArray(graph.Banks)
	// Check if route.FromBIC exists in graph banks
	if !utils.IsInArray(route.FromBIC, banksBICs) {
		return errors.New(`The given route could not be appended to the graph because
		 route's from Bank doesn't match any graph's bank`)
	}

	// Check if route.ToBIC exists in graph banks
	if !utils.IsInArray(route.ToBIC, banksBICs) {
		return errors.New(`The given route could not be appended to the graph because
		 route's to Bank doesn't match any graph's bank`)
	}

	graph.Routes = append(graph.Routes, route)
	return nil
}

// TODO error handlin
func (graph *BankConnectionsGraph) FindRoute(from string, to string) ([]*Bank, error) {
	path := dijkstraAlgorithm(graph, from, to)
	result := make([]*Bank, 0)
	for i := 0; i < len(path); i++ {
		result = append(result, graph.Banks[path[i]])
	}
	return result, nil
}

func dijkstraAlgorithm(targetGraph *BankConnectionsGraph, from string, to string) []string {
	var vertexs map[string]int
	banksBICs := mapKeyToArray(targetGraph.Banks)

	// Set new graph

	graph := dijkstra.NewGraph()
	for i := 0; i < len(banksBICs); i++ {
		vertexs[banksBICs[i]] = i
		graph.AddVertex(i)
	}

	for i := 0; i < len(targetGraph.Routes); i++ {
		graph.AddArc(vertexs[targetGraph.Routes[i].FromBIC], vertexs[targetGraph.Routes[i].ToBIC], 1)
	}

	path, _ := graph.Shortest(vertexs[from], vertexs[to])

	intRoutes := path.Path

	result := make([]string, 0)

	revertedVertexes := utils.ReverseMap(vertexs)

	for i := 0; i < len(intRoutes); i++ {
		result = append(result, revertedVertexes[intRoutes[i]])
	}

	return result
}

func mapKeyToArray(target map[string]*Bank) []string {
	keys := make([]string, 0, len(target))
	for k := range target {
		keys = append(keys, k)
	}
	return keys
}
