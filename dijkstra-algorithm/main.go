package main

import (
	"fmt"
	"math"
)

func FindLowestCostNode(costs map[string]float64, processed map[string]bool) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""

	for node, cost := range costs {
		if cost < lowestCost && !processed[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func dijkstra(graph map[string]map[string]float64, start string) (map[string]float64, map[string]string) {
	costs := make(map[string]float64)
	parents := make(map[string]string)
	processed := make(map[string]bool)

	for node := range graph {
		costs[node] = math.Inf(1)
		parents[node] = ""
	}
	costs[start] = 0

	node := FindLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for neighbor, weight := range neighbors {
			newCost := cost + weight
			if newCost < costs[neighbor] {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}

		processed[node] = true
		node = FindLowestCostNode(costs, processed)
	}

	return costs, parents
}

func main() {
	graph := map[string]map[string]float64{
		"A": {"B": 1, "C": 4},
		"B": {"C": 2, "D": 5},
		"C": {"D": 1},
		"D": {},
	}

	startNode := "A"
	distances, parents := dijkstra(graph, startNode)

	fmt.Println(distances)
	fmt.Println(parents)
}
