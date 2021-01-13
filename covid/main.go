package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
	"os"
	"strconv"
)

var random *rand.Rand // rand for the project

func main() {
	//Declare variables
	var b, numDays, recoveryPeriod int
	var vr, qp, cfr float64
	var numRows, numCols int

	var initialBoard Board
	//capturing the parameters from command line- command, manner of data feeding, filename, two code for neighbourhood type and manner of transition,
	// number of days for running simulation, recovery period, vaccination rate, quarantine probability, case fatality rate

	//filename is name of  the file from which to read and assign the data
	filename := os.Args[2] //"name of the file"
	// b parameter is two digit code to represent the neighbourhood and manner of transition
	//first digit of code- 0 for Moore case and 1 for random (gaussian) case
	//second digit of the code- 0 for determinisitc and 1 for probabilistic

	agr3, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		panic("Error: Problem converting neighbourhood  parameter to an integer number.")
	}
	b = agr3

	// numDays is the number of days (i.e.  number of generations) to simulate the cellular automata

	arg4, err4 := strconv.Atoi(os.Args[4])
	if err4 != nil {
		panic("Error: Problem converting number of generations to an integer.")
	}
	numDays = arg4

	// recoveryPeriod is the number of days for recovery from disease/infection

	arg5, err5 := strconv.Atoi(os.Args[5])
	if err5 != nil {
		panic("Error: Problem converting number of generations to an integer.")
	}
	recoveryPeriod = arg5
	// vr is % of vaccination done per day
	arg6, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic("Error: Problem in parsing from string to float64.")
	}
	vr = arg6
	//qp is quarantine probability- level of restrictions imposed. qp=1 means complete lockdown
	arg7, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic("Error: Problem in parsing from string to float64.")
	}
	qp = arg7
	//cfr is % case fatality rate- fraction of persons dying on a day out of those infected
	arg8, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic("Error: Problem in parsing from string to float64.")
	}
	cfr = arg8

	// parameter that decides whether to generate data for initial board randomly by reading parameters from a file  or by reading persons status from a data file
	if os.Args[1] == "readdata" {
		initialBoard, numRows, numCols = ReadData(filename)

	} else if os.Args[1] == "readparam" {
		initialBoard, numRows, numCols = PopulateData(filename)
	}
	fmt.Printf("row %d, cols %d", numRows, numCols)

	// Set cell width; it should be set in such manner that the pic does not overflow the screen
	cellWidth := 1

	// Save initial image
	fmt.Println("Saving the initial status board (before the simulation) in file initialBoard.png")
	pic1 := DrawBoard(initialBoard, cellWidth)
	pic1.SaveToPNG("initialBoard.png")

	// Generate boards for each day
	fmt.Println("The simulation started")
	boards := SimulateCovidSpread(initialBoard, b, numDays, recoveryPeriod, vr, qp, cfr)
	fmt.Println("The simulation ended")
	//fmt.Println(len(boards))
	WriteSummaryStatusOverTimeToFile(boards, "StatusOverTime.txt")
	finalBoard := boards[(len(boards) - 1)]

	//Draw the GIF
	fmt.Println("creating gif file")
	imageList := DrawBoards(boards, cellWidth)

	gifhelper.Process(imageList, "AnimatedGIF")

	// Save final image
	fmt.Println("Saving the final status board (after the simulation) in file finalBoard.png")
	pic2 := DrawBoard(finalBoard, cellWidth)
	pic2.SaveToPNG("finalBoard.png")

	//Print final game board

	PrintBoard(finalBoard)
	WriteSummaryStatusToFile(finalBoard, "StatusAtTheEnd.txt")

}
