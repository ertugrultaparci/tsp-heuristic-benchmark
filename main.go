package main

import (
	"fmt"
	"tsp-heuristic-benchmark/simmulatedannealing"
)

func main() {

	res := simmulatedannealing.Simmulationannealing("data/data.tsp", 100, 100)
	fmt.Println("Simulation Annealing Result is:", res)
}
