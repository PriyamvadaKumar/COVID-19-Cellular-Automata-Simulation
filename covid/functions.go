package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Set up random number

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// SimulateCovidSpread is a function to generate Boards  for each Day and store the them in a slice of the type Board
func SimulateCovidSpread(initialBoard Board, b, numDays, recoveryPeriod int, vr, qp, cfr float64) []Board {
	boards := make([]Board, numDays+1) // create blank slice of Board
	boards[0] = initialBoard           // store initial Board as the first element of the slice of Board
	for i := 1; i <= numDays; i++ {

		boards[i] = UpdateIntoNewBoard(boards[i-1], b, i, recoveryPeriod, vr, qp, cfr) // generate new board based on  the previous board and the rules of disease transmission

	}
	return boards
}

// NewBoard generates new Blank Board with a given Board
func NewBoard(board Board) Board {
	newBoard := make(Board, len(board))
	for r := range newBoard {
		newBoard[r] = make([]Cell, len(board[r]))
	}
	return newBoard

}

// CopyBoard copies the current for generating a new Board which is a 2D slice, with each element being Cell
func CopyBoard(currentBoard Board) Board {
	newBoard := NewBoard(currentBoard)
	for i := range currentBoard {
		for j := range currentBoard[i] {

			newBoard[i][j] = currentBoard[i][j]

		}
	}
	return newBoard

}

// UpdateIntoNewBoard generates a new board  based on currentBoard and the disease transmission rules
func UpdateIntoNewBoard(currentBoard Board, b, numDays, recoveryPeriod int, vr, qp, cfr float64) Board {
	newBoard := NewBoard(currentBoard) // New board
	for i := range newBoard {
		for j := range newBoard[i] {
			newBoard[i][j] = UpdateCell(currentBoard, i, j, b, numDays, recoveryPeriod, vr, qp, cfr) // update the cell of new  Board

		}
	}

	return newBoard
}

// UpdateCell updates the cell
func UpdateCell(currentBoard Board, r, c, b, numDays, recoveryPeriod int, vr, qp, cfr float64) Cell {
	cell := currentBoard[r][c]
	var n, N int

	// code b is two digits; First digit is for neighbourhood type and second digit represents manner of transition from S to I (deterministic or probabilistic)
	//if first digit is one, represent mobility of the person by Moore "neighbourhood"  and
	//if the first digit is 2, represent mobility of a person by "neighbourhood" sampled from gaussian distribtuion
	// Further if second digit is 1- deterministic transition from S to I if there is even one infected person in the "neighbourhood"
	//Further if second digit is 2- probabilistic transition from S to I on the basis of sampling a probability equal to
	// fraction of total infected persons in the "neibourhood".
	// Further if a person is vaccinated ( modelled with probability vr) , transition from S to V for he has immunity.
	//Further if a person is in infected state but quarantined, then he will not transmit and he should not be counted
	//for spreading infection to others; we model this by a quarantine probability qp to reflect extent to
	//which quarantine rules are enforced. if qp=1, it means full lockdown and qp=0, no restrcitions
	//The transition from I(Infected) to D (Dead) is based on the  case fatality rate( i.e. % death rate among infected),
	// which is modelled by sampling with probability equal to fatality rate.
	if b == 11 || b == 12 {
		n, N = CountInfectedInMooreNeighbourhood(currentBoard, r, c, qp) //count infected persons in Moore neighbourhood
	} else if b == 21 || b == 22 {
		n, N = CountInfectedInRandomNeighbourhood(currentBoard, r, c, qp) //count infected persons in random neighbourhood
	}

	switch currentBoard[r][c].status {
	case "S":
		if random.Float64() < vr/100.0 {
			cell.status = "V" // Transition status from S to V
		} else if (b == 11 || b == 21) && n > 0 { //  deterministic transition: S to I if there is any infected persons(n) around in the "neighbourhood"
			cell.status = "I"     // Transition status from S to I
			cell.dayNum = numDays // store the day of infection
		} else if (b == 12 || b == 22) && (random.Float64() < float64(n)/float64(N)) { // Probabilistic transition-Changing the status from S to I based on probability equal fraction of infected people in the neighbourhood
			cell.status = "I"
			cell.dayNum = numDays
		}
	case "I":
		if (numDays-currentBoard[r][c].dayNum <= recoveryPeriod) && (random.Float64() < cfr/100.0) { //What is the chance of sampling in the zone fatality
			cell.status = "D"
		} else if numDays-currentBoard[r][c].dayNum > recoveryPeriod { // Whether time since the day of infection is more than recovery period in days
			cell.status = "R"

		}

	}

	return cell
}

// CountInfectedInMooreNeighbourhood counts infected in Moore neighbourhood
func CountInfectedInMooreNeighbourhood(board Board, r, c int, qp float64) (int, int) {
	count := 0 // Initialize  the counter for counting infected persons in the neibourhood
	total := 8 // Total persons around the person in the Moore neighbourhood are 8

	//initialize count infected persons in Moore neighbourhood
	for i := r - 1; i <= r+1; i++ { //looping in the neighbor of a cell (r,c)
		for j := c - 1; j <= c+1; j++ {
			if (i != r || j != c) && InField(board, i, j) == true { //exclude the instant cell and remain on the board
				if board[i][j].status == "I" {
					count++ // increment the counter when person with status I i.e. infected is found in the neigbourhood
				}
			}

		}
	}
	count = int(float64(count) * qp)

	return count, total

}

// CountInfectedInRandomNeighbourhood counts infected in Random neighbourhood
func CountInfectedInRandomNeighbourhood(board Board, r, c int, qp float64) (int, int) {
	count := 0 // Initialize the counter for counting infected persons in the neibourhood
	numRows := len(board)
	numCols := len(board[0])

	//Range of random neighbourhood
	minRow, maxRow := RandomRowRange(numRows, r)
	minCol, maxCol := RandomRowRange(numCols, c)
	//initialize count infected persons in Random neighbourhood
	for i := minRow; i <= maxRow; i++ { //looping in the neighbor of a cell (r,c)
		for j := minCol; j <= maxCol; j++ {
			if (i != r || j != c) && InField(board, i, j) == true { //exclude the instant cell and remain on the board
				if board[i][j].status == "I" {
					count++ // increment the counter when person with status I i.e. infected is found in the neigbourhood
				}
			}

		}
	}
	total := (maxRow-minRow+1)*(maxCol-minCol+1) - 1 // total persons around inside the random neighbourhood
	count = int(float64(count) * qp)

	return count, total

}

//RandomRowRange return minRow and maxRow of the region choosen randomly
func RandomRowRange(numRows, row int) (minRow int, maxRow int) {
	row1 := int(rand.NormFloat64()*float64(numRows)/float64(2) + float64(row))
	row1 = row1 % numRows
	if row1 < 0 {
		row1 = row1 + numRows
	}
	row2 := int(rand.NormFloat64()*float64(numRows)/float64(2) + float64(row))
	row2 = row2 % numRows
	if row2 < 0 {
		row2 = row2 + numRows
	}

	if row1 < row2 {
		minRow = row1
		maxRow = row2

	} else {
		minRow = row2
		maxRow = row1

	}
	return

}

//RandomColRange return minCol and maxCol of the region choosen randomly
func RandomColRange(numCols, col int) (minCol int, maxCol int) {
	col1 := int(rand.NormFloat64()*float64(numCols)/float64(2) + float64(col))
	col1 = col1 % numCols
	if col1 < 0 {
		col1 = col1 + numCols
	}
	col2 := int(rand.NormFloat64()*float64(numCols)/float64(2) + float64(col))
	col2 = col2 % numCols
	if col2 < 0 {
		col2 = col2 + numCols
	}

	if col1 < col2 {
		minCol = col1
		maxCol = col2

	} else {
		minCol = col2
		maxCol = col1

	}
	return

}

// InField function is to check whether we are on the board, given the value of row and column
func InField(Board Board, r, c int) bool {
	numRows := len(Board)
	numCols := len(Board[0])
	if r < 0 || r >= numRows {
		return false
	}
	if c < 0 || c >= numCols {
		return false
	}
	return true
}

// NewStatusBoard generates a new 2-D string array( board) for status
func NewStatusBoard(numRows, numCols int) [][]string {
	newBoard := make([][]string, numRows)
	for r := range newBoard {
		newBoard[r] = make([]string, numCols)
	}
	return newBoard

}

//BoardToStatusBoard returns a 2-D string array( board) with status of persons as cell, given the Board
func BoardToStatusBoard(board Board) [][]string {
	statusBoard := NewStatusBoard(len(board), len(board[0]))
	for i := range statusBoard {
		for j := range statusBoard[i] {

			statusBoard[i][j] = board[i][j].status
		}
	}
	return statusBoard
}

//PrintBoard prints the status of the cells on the board
func PrintBoard(board Board) {
	for i := range board {
		for j := range board[0] {
			fmt.Printf("%s", board[i][j].status)

		}
		fmt.Println()

	}
}
