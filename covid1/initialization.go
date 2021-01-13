package main

import (
	"math/rand"
)

//InitializeClasses populates the slice that has classes( population proportions at various states/time) for a region
// It sets entire population as susceptible i.e. S=1.0
func InitializeClasses(numStates int) []float64 {
	s := make([]float64, numStates)

	s[0] = 1.0

	return s
}

//InitializeReciprocity  generates reciprocity parameters to represent the  interations between the regions/Cells
func InitializeReciprocity(numRegions int) []float64 {
	c := make([]float64, numRegions)
	for i := range c {
		c[i] = rand.Float64()
	}

	return c
}

//InitializeBoard initialises the board
func InitializeBoard(numRegions, numStates int) *Board {

	var b Board
	b.cells = make([]*Cell, numRegions)

	for i := range b.cells {
		var cell Cell
		//set the slice storing class proportions for the length of nummber of states
		cell.classProportions = make([]float64, numStates)
		//Initialse  slice with proportion of susceptible population as 1
		p := InitializeClasses(numStates)
		copy(cell.classProportions, p)
		//Initialise population and area
		cell.population = rand.Intn(1000000) + 100000
		cell.area = rand.Float64()*1000 + 100.0
		//set the slice storing reciprocity for the length of number of regions
		cell.reciprocity = make([]float64, numRegions)

		//Initialse  slice with proportion of susceptible population as 1
		r := InitializeReciprocity(numRegions)
		copy(cell.reciprocity, r)
		//Assign the cell populated to the board
		b.cells[i] = &cell

	}
	// Choose a region to introduce infection
	j := rand.Intn(numRegions)
	//Choose proportion infected
	p := 0.1
	b.cells[j].classProportions[1] = p
	b.cells[j].classProportions[0] = 1 - p
	return &b
}

//Lambdas generates  lambda for all the states
func Lambdas(numStates int) []float64 {
	lambda := make([]float64, numStates)
	for i := range lambda {
		lambda[i] = 0.8
	}
	lambda[0] = 0.0

	return lambda
}

//Gammas generates gamma for all the states
func Gammas(numStates int) []float64 {
	gamma := make([]float64, numStates)
	for i := range gamma {
		gamma[i] = 0.1
	}
	gamma[0] = 0.1

	return gamma
}
