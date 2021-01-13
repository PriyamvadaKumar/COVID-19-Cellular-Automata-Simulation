package main

//Board declared
type Board struct {
	cells []*Cell
}

// Cell holds proportion of persons of various categories/states  and other data of the region
type Cell struct {
	classProportions []float64 //to store proportion of persons in different states in a cell(region)

	population  int       // to store the population of   a region represented by the cell
	area        float64   // to store area of the region represented by the cell
	reciprocity []float64 // to store reciprocity parameters to account for interaction between the  regions

}

//Periods contain the periods for infection etc
type Periods struct {
	Ti int
	Tp int
	Tl int
	Tr int
}
