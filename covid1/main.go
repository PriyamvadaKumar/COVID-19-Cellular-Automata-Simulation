package main

import (
	"fmt"
	"gifhelper"
)

func main() {
	//Declare variables
	// Variable for number regions which are represented by the cells on the board
	var numRegions int = 5
	//The number of Days for whcih the simulation would be done
	var numDays int = 100

	//Declare and set the time periods for various phases of illeness etc
	var statePeriods Periods
	statePeriods.Ti = 2 // incubation period ( when a perosn is infectious but not showing syptoms of illness)

	statePeriods.Tp = 2 // proper infection period ( when a person is infectious as well as showing symptoms)

	statePeriods.Tl = 2 // latency period ( When a person is not infectious but showing symptoms of illness)

	statePeriods.Tr = 7 // immunity period ( neither infectious nor showing symptoms i.e. recovered, a period after which immunity vanishes)
	//Each day of the various phases corresponds to a state
	numStates := statePeriods.Ti + statePeriods.Tp + statePeriods.Tl + statePeriods.Tr + 1 // total number of states
	numStatesSI := statePeriods.Ti + statePeriods.Tp + statePeriods.Tl                     // number of states till the end of infection period
	//Cell width in terms of pixels
	cellWidth := 30
	canvasWidth := 1000
	//Initialise the Board(This is 1-Dimensional Board)
	initialBoard := InitializeBoard(numRegions, numStates)

	PrintBoard(initialBoard, numStatesSI)
	pic1 := DrawBoard(initialBoard, canvasWidth, cellWidth, numStatesSI)
	pic1.SaveToPNG("initialBoardCovidPaper.png")

	lambda := Lambdas(numStates)
	fmt.Println(lambda)
	gamma := Gammas(numStates)

	fmt.Println(gamma)
	// Generate boards for each day
	fmt.Println("The simulation started")
	boards := SimulateCovidSpread(initialBoard, statePeriods, lambda, gamma, numDays)
	fmt.Println("The simulation ended")

	//Print the data of boards over time

	fmt.Println("Printing S, I and R for each rigion for each day")
	PrintBoards(boards, numStatesSI)
	//Write the data of boards over time( region wise)
	WriteStatusOverTimeToFile(boards, numStatesSI, "StatusOverTime.txt")
	//Write the data of boards over time( region wise)
	WriteSummaryStatusOverTimeToFile(boards, numStatesSI, "SummaryStatusOverTime.txt")

	//Print the data of final board

	PrintBoard(boards[numDays], numStatesSI)

	imageList := DrawBoards(boards, canvasWidth, cellWidth, numStatesSI)
	gifhelper.Process(imageList, "AnimatedCovidPaperGIF")
	pic2 := DrawBoard(boards[numDays], canvasWidth, cellWidth, numStatesSI)
	pic2.SaveToPNG("finalBoardCovidPaper.png")

}
