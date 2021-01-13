package main

import "fmt"

// SimulateCovidSpread is a function to generate Boards  for each Day and store the them in a slice of the type Board
func SimulateCovidSpread(initialBoard *Board, statePeriods Periods, lambda, gamma []float64, numDays int) []*Board {
	boards := make([]*Board, numDays+1) // create blank slice of Board

	boards[0] = initialBoard // store initial Board as the first element of the slice of Board

	for i := 1; i <= numDays; i++ {

		boards[i] = UpdateBoard(boards[i-1], statePeriods, lambda, gamma) // generate new board based on  the previous board  and rules of disease spread

	}
	return boards
}

// UpdateBoard generates a new board  based on the disease transmission rules
func UpdateBoard(currentBoard *Board, statePeriods Periods, lambda, gamma []float64) *Board {
	newBoard := CopyBoard(currentBoard)

	for i, cell := range newBoard.cells {
		cell.UpdateCell(currentBoard, statePeriods, i, lambda, gamma)
	}

	return newBoard
}

// UpdateCell updates  a cell(region)
func (cell *Cell) UpdateCell(currentBoard *Board, statePeriods Periods, i int, lambda, gamma []float64) {
	cell.UpdateCellS(currentBoard, statePeriods, i, lambda)
	cell.UpdateCellI(currentBoard, statePeriods, i, lambda, gamma)
	cell.UpdateCellR(currentBoard, statePeriods, i, gamma)

}

// UpdateCellS updates the % susceptibles in a cell(region)
func (cell *Cell) UpdateCellS(currentBoard *Board, statePeriods Periods, i int, lambda []float64) {
	term1 := currentBoard.cells[i].classProportions[0]
	T := statePeriods.Ti + statePeriods.Tp + statePeriods.Tl + statePeriods.Tr
	term2 := currentBoard.cells[i].classProportions[T]

	cell.classProportions[0] = term1 + term2 - InternalTerm(currentBoard, statePeriods, i, lambda) - ExternalTerm(currentBoard, statePeriods, i, lambda)

}

// UpdateCellI updates the % infected persons in a  cell(region)
func (cell *Cell) UpdateCellI(currentBoard *Board, statePeriods Periods, i int, lambda, gamma []float64) {
	cell.classProportions[1] = InternalTerm(currentBoard, statePeriods, i, lambda) + ExternalTerm(currentBoard, statePeriods, i, lambda)
	T := statePeriods.Ti + statePeriods.Tp + statePeriods.Ti

	for q := 2; q <= T; q++ {
		cell.classProportions[q] = (1 - gamma[q-1]) * currentBoard.cells[i].classProportions[q-1]
	}

}

// UpdateCellR updates the % recovered persons  in a  cell(region)
func (cell *Cell) UpdateCellR(currentBoard *Board, statePeriods Periods, i int, gamma []float64) {
	T := statePeriods.Ti + statePeriods.Tp + statePeriods.Ti

	cell.classProportions[T+1] = currentBoard.cells[i].classProportions[T] + InfectiveCuredTerm(currentBoard, statePeriods, i, gamma)
	for q := T + 2; q <= T+statePeriods.Tr; q++ {
		cell.classProportions[q] = currentBoard.cells[i].classProportions[q-1]
	}

}

//InternalTerm calculates term in updation of % Susceptible that relates to infection  from within the region
func InternalTerm(currentBoard *Board, statePeriods Periods, i int, lambda []float64) float64 {
	var v float64
	T := statePeriods.Ti + statePeriods.Tp + statePeriods.Ti

	roi := currentBoard.cells[i].classProportions[0] * float64(currentBoard.cells[i].population) / (currentBoard.cells[i].area * 1000.0)

	for q := 1; q <= T; q++ {
		v = v + lambda[q]*roi*currentBoard.cells[i].classProportions[0]*currentBoard.cells[i].classProportions[q]

	}
	return v

}

//ExternalTerm calculates term in updation of proportion of Susceptible infected due others regions
func ExternalTerm(currentBoard *Board, statePeriods Periods, i int, lambda []float64) float64 {
	var v float64
	T := statePeriods.Ti + statePeriods.Tp + statePeriods.Ti

	roi := currentBoard.cells[i].classProportions[0] * float64(currentBoard.cells[i].population) / (currentBoard.cells[i].area * 1000.0)
	Ci := currentBoard.cells[i].reciprocity

	for j := 0; j < len(currentBoard.cells); j++ {

		for q := 1; q <= T; q++ {
			v = v + Ci[j]*lambda[q]*roi*currentBoard.cells[i].classProportions[0]*currentBoard.cells[j].classProportions[q]
		}

	}
	return v

}

//InfectiveCuredTerm calculates the term in updation of R that relates to cured persons
func InfectiveCuredTerm(currentBoard *Board, statePeriods Periods, i int, gamma []float64) float64 {
	T := statePeriods.Ti + statePeriods.Tp + statePeriods.Ti
	var v float64
	for q := 1; q <= T-1; q++ {
		v = gamma[q] * currentBoard.cells[i].classProportions[q]

	}
	return v
}

// CopyBoard copies the current board
func CopyBoard(currentBoard *Board) *Board {
	var newBoard Board
	newBoard.cells = make([]*Cell, len(currentBoard.cells))
	for i := range newBoard.cells {
		newBoard.cells[i] = CopyCell(currentBoard.cells[i])
	}

	return &newBoard
}

// CopyCell copies the current cell
func CopyCell(currentCell *Cell) *Cell {

	var newCell Cell
	newCell.classProportions = make([]float64, len(currentCell.classProportions))
	newCell.reciprocity = make([]float64, len(currentCell.reciprocity))

	copy(newCell.classProportions, currentCell.classProportions)

	newCell.population = currentCell.population

	newCell.area = currentCell.area
	copy(newCell.reciprocity, currentCell.reciprocity)

	return &newCell

}

//PrintBoards prints S, I and R of all regions(Cells) on the board for all days- one board for each day
// q is total states till the end of infection period- one state for each day. q= Ti+Tp+Tl
func PrintBoards(boards []*Board, q int) {
	for i := range boards {
		PrintBoard(boards[i], q)
		fmt.Printf("%v the day\n", i)
	}
}

//PrintBoard prints cells on a Board
func PrintBoard(board *Board, q int) {

	for _, cell := range board.cells {

		fmt.Printf("%.2f ", cell.classProportions[0])
		fmt.Printf("%.2f ", SumSlice(cell.classProportions[1:q+1]))
		fmt.Printf("%.2f ", SumSlice(cell.classProportions[q+1:]))
		fmt.Printf("\n")
	}

}

//ClassSum retuns the percentage  of total persons in different classProportions namely S, I and R for each region/cell
//It takes cell and q=Ti+Tp+Tl,  time/days  when the infection phase ends
func ClassSum(cell *Cell, q int) (float64, float64, float64) {
	var s, i, r float64
	s = cell.classProportions[0]
	i = SumSlice(cell.classProportions[1 : q+1])
	r = SumSlice(cell.classProportions[q+1:])

	return s, i, r

}

//SumSlice determines the sum of elements in a slice
func SumSlice(a []float64) float64 {
	var v float64
	for i := range a {
		v = v + a[i]
	}
	return v
}
