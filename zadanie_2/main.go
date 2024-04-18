package main

import (
	"time"
)

func main() {
	sim := NewSimulation(50, 25)
	sim.Begin()
	for {
		sim.Update()
		time.Sleep(500 * time.Millisecond)
	}
}