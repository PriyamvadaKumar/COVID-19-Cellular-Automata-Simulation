package main

// Cell type is to hold status and day of infection
type Cell struct {
	status string //to store "S"/"V"/"I"/"R"/"D" for status of inividual- susceptible/vaccinated/infected/recovered/dead
	dayNum int    //to store the day of infection
}

//Board is a 2D slice of Cells which are struct type to hold status and the day of infection
type Board [][]Cell
