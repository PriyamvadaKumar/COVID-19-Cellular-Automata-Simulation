package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//ReadData reads the given  file, and then returns board, rows , columns
func ReadData(filename string) (board Board, numRows int, numCols int) {

	f, err3 := os.Open(filename)

	if err3 != nil {
		fmt.Println("Error: something went wrong in opening the file")
		os.Exit(1)
	}

	defer f.Close()

	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Error: some error during the file reading")
		os.Exit(1)
	}
	var p []string = strings.Split(lines[0], " ")
	numRows, err4 := strconv.Atoi(p[0])
	if err4 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	numCols, err5 := strconv.Atoi(p[1])
	if err5 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}

	board = make([][]Cell, numRows)
	for i := 0; i < numRows; i++ {
		var temp []string = strings.Split(lines[i+1], "")
		fmt.Println(temp)

		for j := 0; j < numCols; j++ {
			var cell Cell

			cell.status = temp[j]
			board[i] = append(board[i], cell)
		}

	}

	return
}

//PopulateData reads the given  file having one number of rows, number of columns and number of persons infected at the start
//in one line and then returns a Board  randomly populated with given number of infected persons, along with rows , columns
func PopulateData(filename string) (board Board, numRows int, numCols int) {
	f, err3 := os.Open(filename)

	if err3 != nil {
		fmt.Println("Error: something went wrong in opening the file")
		os.Exit(1)
	}

	defer f.Close()

	var line string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
	}
	if scanner.Err() != nil {
		fmt.Println("Error: some error during the file reading")
		os.Exit(1)
	}
	var p []string = strings.Split(line, " ")
	numRows, err4 := strconv.Atoi(p[0])
	if err4 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	numCols, err5 := strconv.Atoi(p[1])
	if err5 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}

	numInfectedStart, err5 := strconv.Atoi(p[1]) // number of persons to be taken as infected at the begining of simulation
	if err5 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	// Populate the board with cells(persons) with status "S" as susceptible
	board = make([][]Cell, numRows)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			var cell Cell
			cell.status = "S"
			cell.dayNum = 0

			board[i] = append(board[i], cell)
		}

	}
	// Infect some perons at the start randomly
	for n := 0; n < numInfectedStart; n++ {

		r := rand.Intn(numRows) // generate a random integer betweem [0, numRows) to select a row
		c := rand.Intn(numCols) //generate a random integer betweem [0, numCols) to select a column

		board[r][c].status = "I"
	}

	return
}

//WriteBoardToFile takes a board and a filename as a string and writes the status of board into the file
func WriteBoardToFile(board Board, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for r := range board {
		for c := range board[r] {
			fmt.Fprint(file, board[r][c].status)
		}
		fmt.Fprintln(file)
	}
}

//WriteSummaryStatusOverTimeToFile takes slice of boards and a filename and writes the summary of status over time into the file
func WriteSummaryStatusOverTimeToFile(boards []Board, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for t := 0; t < len(boards); t++ {
		countS := 0
		countV := 0
		countI := 0
		countR := 0
		countD := 0
		board := boards[t]
		for r := range board {
			for c := range board[r] {

				switch board[r][c].status {

				case "S":
					countS++
				case "V":
					countV++
				case "I":
					countI++
				case "R":
					countR++
				case "D":
					countD++
				}

				//fmt.Fprint(file, c, countS, countI, countR, countD)
				//fmt.Fprintln(file)

			}

		}

		fmt.Fprint(file, countS, countV, countI, countR, countD)
		fmt.Fprintln(file)
	}

}

//WriteSummaryStatusToFile takes a board and a filename as a string and writes the summary of status into the file
func WriteSummaryStatusToFile(board Board, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	var countS, countV, countI, countR, countD int

	for r := range board {
		for c := range board[r] {

			switch board[r][c].status {

			case "S":
				countS++
			case "V":
				countV++
			case "I":
				countI++
			case "R":
				countR++
			case "D":
				countD++
			}

		}

	}

	fmt.Fprint(file, "No of Susceptible persons: ", countS, "\n", "No of Vaccinated persons: ", countV, "\n", "No of Infected persons: ", countI, "\n", "No of Recovered persons: ", countR, "\n", "No of Dead persons: ", countD)
	fmt.Fprintln(file)

}
