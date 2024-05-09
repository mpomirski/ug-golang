package main

import (
	"fmt"
	"math/rand"
)

type FieldType uint8
const (
	Empty FieldType = iota
	Tree
	OldTree
	BurningTree
	BurnedTree
	Lightning
)

type Forest struct {
	plane [][]FieldType
	initialTrees int
}
func (forest *Forest) DrawForest() {
	for _, row := range forest.plane {
		for _, cell := range row {
			switch cell {
			case Empty:
				fmt.Print(" ")
			case Tree:
				fmt.Print("T")
			case BurningTree:
				fmt.Print("🔥")
			case BurnedTree:
				fmt.Print("↟")
			case OldTree:
				fmt.Print("𖠰")
			case Lightning:
				fmt.Print("⚡")
			}
		}
		fmt.Println()
	}
}


func (forest *Forest) GetBurntCount() int {
	burnt := 0
	for i := 0; i < len(forest.plane); i++ {
		for j := 0; j < len(forest.plane[i]); j++ {
			if forest.plane[i][j] == BurnedTree {
				burnt++
			}
		}
	}
	return burnt
}

func (forest *Forest) GetInitialTreeCount() int {
	trees := 0
	for i := 0; i < len(forest.plane); i++ {
		for j := 0; j < len(forest.plane[i]); j++ {
			if forest.plane[i][j] == Tree {
				trees++
			}
		}
	}
	return trees
}

func (forest *Forest) GetForestSize() int {
	return len(forest.plane) * len(forest.plane[0])
}

func NewForest(sizeX, sizeY, initialTrees int) *Forest {
	forest := &Forest{
		plane: make([][]FieldType, sizeY),
		initialTrees: initialTrees,
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
	count := 0
	for count < forest.initialTrees {
		for i := 0; i < len(forest.plane); i++ {
			for j := 0; j < len(forest.plane[i]); j++ {
				if rand.Intn(100) < 10 && count < forest.initialTrees{
					forest.SetTree(i, j)
					count++
				}
				if rand.Intn(100) < 1 && count < forest.initialTrees{
					forest.SetOldTree(i, j)
					count++
				}
			}
		}
	}
}
func (forest *Forest) SetLightning(x, y int) {
	forest.plane[x][y] = Lightning
}

func (forest *Forest) SetOldTree(x, y int) {
	forest.plane[x][y] = OldTree
}

func (forest *Forest) RandomLightning(){
	x := rand.Intn(len(forest.plane))
	y := rand.Intn(len(forest.plane[0]))
	if forest.plane[x][y] == OldTree {
		forest.SetLightning(x, y)
	}
	if forest.plane[x][y] == Tree && rand.Intn(100) < 1 {
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
					if forest.plane[neighbor[0]][neighbor[1]] == OldTree {
						burningTrees = append(burningTrees, [2]int{neighbor[0], neighbor[1]})
					}
				}
			case Lightning:
				forest.SetBurningTree(i, j)
			case Tree:
				if rand.Intn(100) < 10 {
					forest.SetOldTree(i, j)
				}
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
