package main

import (
	"time"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("Initial trees:", i * 10)
		initialTrees:= i * 10
		sim := NewSimulation(50, 25, initialTrees)
		sim.DisableDrawing()
		sim.Begin()
		steps := 30
		for i := 0; i < steps; i++{
			sim.Update()
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Printf("Burnt (%% of forest): %d %.2f%%\n", sim.forest.GetBurntCount(), float64(sim.forest.GetBurntCount())*100.0 / float64(sim.forest.GetForestSize()))
	}
	// initialTrees:= 500
	// fmt.Printf("Initial trees: %d\n", initialTrees)
	// sim := NewSimulation(50, 25, initialTrees)
	// sim.Begin()
	// steps := 50
	// for i := 0; i < steps; i++{
	// 	sim.Update()
	// 	time.Sleep(200 * time.Millisecond)
	// }
}