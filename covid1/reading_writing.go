package main

import (
	"fmt"
	"log"
	"os"
)

/*
//ReadData reads the given  file, and then returns board, rows , columns
func ReadData(filename string) (initialBoard Board) {

	f, err2 := os.Open(filename)

	if err2 != nil {
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
	var l0 []string = strings.Split(lines[0], " ")
	numRegions, err3 := strconv.Atoi(l0[0])
	if err3 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	numDays, err4 := strconv.Atoi(l0[1])
	if err4 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	var l1 []string = strings.Split(lines[1], " ")
	Ti, err5 := strconv.Atoi(l1[0])
	if err5 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	Tp, err6 := strconv.Atoi(l1[1])
	if err6 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	Tl, err7 := strconv.Atoi(l1[2])
	if err7 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	Tr, err8 := strconv.Atoi(l1[3])
	if err8 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	var l2 []string = strings.Split(lines[2], " ")
	lambda := make([]float64, 0)
	for i := range l2 {
		var v float64
		v, err9 := strconv.ParseFloat(l2[i], 64)
		if err9 != nil {
			fmt.Println("Error: something went wrong in conversion from string to integer")
		}
		lambda = append(lambda, v)

	}
	var l3 []string = strings.Split(lines[3], " ")
	gamma := make([]float64, 0)
	for i := range l3 {
		var v float64
		v, err9 := strconv.ParseFloat(l3[i], 64)
		if err9 != nil {
			fmt.Println("Error: something went wrong in conversion from string to integer")
		}
		gamma = append(lambda, v)

	}

	var l4 []string = strings.Split(lines[4], " ")
	S, err5 := strconv.ParseFloat(l4[0], 64)
	if err5 != nil {
		fmt.Println("Error: something went wrong in conversion from string to integer")

	}
	var l5 []string = strings.Split(lines[5], " ")
	population := make([]int, 0)
	for i := range l5 {

		v, err9 := strconv.Atoi(l5[i])
		if err9 != nil {
			fmt.Println("Error: something went wrong in conversion from string to integer")
		}
		population = append(population, v)

	}

	var l6 []string = strings.Split(lines[6], " ")
	area := make([]float64, 0)
	for i := range l6 {

		v, err9 := strconv.ParseFloat(l6[i], 64)
		if err9 != nil {
			fmt.Println("Error: something went wrong in conversion from string to integer")
		}
		area = append(area, v)

	}

	C:=make([]float, numRegions+1)
	C[0]=make(float64, numRegions+1)
	for j := 0; j < numRegions; i++ {
		var temp []string = strings.Split(lines[7+i], " ")
		r:=make(float,0)
		for i := range temp {

			e, err9 := strconv.ParseFloat(temp[i], 64)
			if err9 != nil {
				fmt.Println("Error: something went wrong in conversion from string to integer")
			}
			r = append(r, e)

			}
			C = append(C, r)
		}

	}


}


*/
//WriteBoardToFile takes a board and a filename as a string and writes the status of board into the file
func WriteBoardToFile(board *Board, q int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, cell := range board.cells {
		fmt.Fprint(file, cell.classProportions[0], SumSlice(cell.classProportions[1:q+1]), SumSlice(cell.classProportions[q+1:]))
		fmt.Fprintln(file)
	}

}

//WriteStatusOverTimeToFile takes slice of boards,q the states up to the end of infection period and a filename and writes the status of each region over time into the file
func WriteStatusOverTimeToFile(boards []*Board, q int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for t := range boards {
		for _, cell := range boards[t].cells {

			fmt.Fprint(file, cell.classProportions[0], SumSlice(cell.classProportions[1:q+1]), SumSlice(cell.classProportions[q+1:]))
		}
		fmt.Fprintln(file)
	}

}

//WriteSummaryStatusOverTimeToFile takes a board, q states up to the end of infection period and a filename as a string and writes the summary of status over time into the file
func WriteSummaryStatusOverTimeToFile(boards []*Board, q int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	var S, I, R float64
	for t := range boards {
		for _, cell := range boards[t].cells {
			S += cell.classProportions[0] * float64(cell.population)
			I += SumSlice(cell.classProportions[1:q+1]) * float64(cell.population)
			R += SumSlice(cell.classProportions[q+1:]) * float64(cell.population)
		}
		fmt.Fprint(file, S, I, R)
		fmt.Fprintln(file)
	}

}
