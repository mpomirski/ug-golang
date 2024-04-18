package main

import (
	"fmt"
	"math/rand"
)

type FieldType uint8
const (
	Empty FieldType = iota
	Tree
	BurningTree
	BurnedTree
	Lightning
)

type Forest struct {
	plane [][]FieldType
}
func (forest *Forest) DrawForest() {
	for _, row := range forest.plane {
		for _, cell := range row {
			switch cell {
			case Empty:
				fmt.Print(" ")
			case Tree:
				fmt.Print("𖠰")
			case BurningTree:
				fmt.Print("🔥")
			case BurnedTree:
				fmt.Print("↟")
			case Lightning:
				fmt.Print("⚡")
			}
		}
		fmt.Println()
	}
}
func NewForest(sizeX, sizeY int) *Forest {
	forest := &Forest{
		plane: make([][]FieldType, sizeY),
	}
	for i := range forest.plane {
		forest.plane[i] = make([]FieldType, sizeX)
	}
	return forest
}
func (forest *Forest) SetTree(x, y int) {
	forest.plane[x][y] = Tree
}
func (forest *Forest) SetBurningTree(x, y int) {
	forest.plane[x][y] = BurningTree
}
func (forest *Forest) SetBurnedTree(x, y int) {
	forest.plane[x][y] = BurnedTree
}
func (forest *Forest) RandomTrees() {
	for i := 0; i < len(forest.plane); i++ {
		for j := 0; j < len(forest.plane[i]); j++ {
			if rand.Intn(100) < 30{
				forest.SetTree(i, j)
			}
		}
	}
}
func (forest *Forest) SetLightning(x, y int) {
	forest.plane[x][y] = Lightning
}

func (forest *Forest) RandomLightning(){
	x := rand.Intn(len(forest.plane))
	y := rand.Intn(len(forest.plane[0]))
	if forest.plane[x][y] == Tree {
		forest.SetLightning(x, y)
	}
}

func (forest *Forest) Update() {
	burningTrees := [][2]int{}
	for i := 0; i < len(forest.plane); i++ {
		for j := 0; j < len(forest.plane[i]); j++ {
			switch forest.plane[i][j] {
			case BurningTree:
				forest.SetBurnedTree(i, j)
				neighbors := forest.GetNeighbors(i, j)
				for _, neighbor := range neighbors {
					if forest.plane[neighbor[0]][neighbor[1]] == Tree {
						burningTrees = append(burningTrees, [2]int{neighbor[0], neighbor[1]})
					}
				}
			case Lightning:
				forest.SetBurningTree(i, j)
			}
		}
	}
	for _, burningTree := range burningTrees {
		forest.SetBurningTree(burningTree[0], burningTree[1])
	}
}

func (forest *Forest) GetNeighbors(x, y int) [][2]int {
	neighbors := [][2]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
	for neighborIndex, neighbor := range neighbors {
		if neighbor[0] < 0 || neighbor[0] >= len(forest.plane) || neighbor[1] < 0 || neighbor[1] >= len(forest.plane[0]) {
			neighbors = append(neighbors[:neighborIndex], neighbors[neighborIndex+1:]...)
		}
	}
	return neighbors
}