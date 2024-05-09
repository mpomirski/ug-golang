package main

import (
	"fmt"
)

type Simulation struct {
	forest *Forest
	step int
	initialTrees int
	drawing bool
}

func NewSimulation(sizeX, sizeY, initialTrees int) *Simulation {
	return &Simulation{
		forest: NewForest(sizeX, sizeY, initialTrees),
		step: 0,
		initialTrees: initialTrees,
		drawing: true,
	}
}

func (sim *Simulation) Begin() {
	sim.forest.RandomTrees()
}

func (sim *Simulation) DisableDrawing() {
	sim.drawing = false
}

func (sim *Simulation) Update() {
	sim.forest.Update()
	sim.forest.RandomLightning()
	if sim.drawing {
		fmt.Print("\033[H\033[2J")
		sim.forest.DrawForest()
		fmt.Println("Step:", sim.step)
		fmt.Printf("Burnt (%% of forest): %d %.2f%%\n", sim.forest.GetBurntCount(), float64(sim.forest.GetBurntCount())*100.0 / float64(sim.forest.GetForestSize()))
		fmt.Printf("Initial trees (%% of forest): %d %.2f%%\n", sim.initialTrees, float64(sim.initialTrees)*100.0/float64(sim.forest.GetForestSize()))
	}
	sim.step++
}