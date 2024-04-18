package main

import (
	"fmt"
)

type Simulation struct {
	forest *Forest
	step int
}

func NewSimulation(sizeX, sizeY int) *Simulation {
	return &Simulation{
		forest: NewForest(sizeX, sizeY),
	}
}

func (sim *Simulation) Begin() {
	sim.forest.RandomTrees()
}

func (sim *Simulation) Update() {
	fmt.Print("\033[H\033[2J")
	sim.forest.Update()
	sim.forest.RandomLightning()
	sim.forest.DrawForest()
	fmt.Println("Step:", sim.step)
	sim.step++
}