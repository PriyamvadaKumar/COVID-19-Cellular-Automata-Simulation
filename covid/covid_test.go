package main

import (
	"reflect"
	"testing"
)

// Testing InField function
//Create type testpair
type InfieldInput struct {
	b Board
	r int
	c int
}

type infieldtestpair struct {
	input InfieldInput
	value bool
}

// Initialise test pairs
var inFieldTests = []infieldtestpair{
	{InfieldInput{Board{{{"S", 0}}}, 0, 0}, true},
	{InfieldInput{Board{{{"S", 0}}}, 1, 1}, false},
	{InfieldInput{Board{{{"S", 0}, {"S", 0}}}, 1, 1}, false},
	{InfieldInput{Board{{{"S", 0}, {"S", 0}}, {{"S", 0}, {"S", 0}}}, 1, 1}, true},
	{InfieldInput{Board{{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}, 1, 1}, true},
	{InfieldInput{Board{{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}, 1, 3}, false},
}

func Test_Infield(t *testing.T) {

	// compare what is expected and what we get
	for _, pair := range inFieldTests {

		v := InField(pair.input.b, pair.input.r, pair.input.c)
		if v != pair.value {
			t.Error("For", pair.input,
				"Expected", pair.value,
				"got", v)
		}
	}

}

// Testing  CountInfectedInMooreNeighbourhood function
//Create type testpair
type CountInfectedInMooreNeighbourhoodInput struct {
	b Board
	r int
	c int
}

type countInfectedInMooreNeighbourhoodtestpair struct {
	input CountInfectedInMooreNeighbourhoodInput
	value int
}

var countInfectedInMooreNeighbourhoodTests = []countInfectedInMooreNeighbourhoodtestpair{

	{CountInfectedInMooreNeighbourhoodInput{Board{{{"S", 0}, {"I", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"I", 0}, {"I", 0}}}, 1, 1}, 3},
	{CountInfectedInMooreNeighbourhoodInput{Board{{{"S", 0}, {"I", 0}, {"S", 0}},
		{{"S", 0}, {"I", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}, 1, 1}, 1},
	{CountInfectedInMooreNeighbourhoodInput{Board{{{"I", 0}, {"I", 0}, {"S", 0}},
		{{"S", 0}, {"I", 0}, {"S", 0}},
		{{"I", 0}, {"I", 0}, {"S", 0}}}, 1, 1}, 4},
	{CountInfectedInMooreNeighbourhoodInput{Board{{{"I", 0}, {"I", 0}, {"I", 0}},
		{{"I", 0}, {"I", 0}, {"I", 0}},
		{{"I", 0}, {"I", 0}, {"I", 0}}}, 1, 1}, 8},
	{CountInfectedInMooreNeighbourhoodInput{Board{{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}, 1, 1}, 0},
}

func Test_CountInfectedInMooreNeighbourhood(t *testing.T) {

	// compare what is expected and what we get
	for _, pair := range countInfectedInMooreNeighbourhoodTests {

		v := CountInfectedInMooreNeighbourhood(pair.input.b, pair.input.r, pair.input.c)

		if v != pair.value {
			t.Error("For", pair.input,
				"Expected", pair.value,
				"got", v)
		}
	}

}

// Testing  UpdateCell function
//Create type testpair
type UpdateCellInput struct {
	b       Board
	r       int
	c       int
	code    int
	numDays int
}

type updatecellinputpair struct {
	input UpdateCellInput
	value Cell
}

var updateCellTests = []updatecellinputpair{
	{UpdateCellInput{Board{{{"S", 0}, {"I", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"I", 0}, {"I", 0}}}, 1, 1, 11, 10}, Cell{"I", 10}},
	{UpdateCellInput{Board{{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}, 1, 1, 11, 20}, Cell{"S", 0}}}

func Test_UpdateCell(t *testing.T) {

	// Comparison between the expected and what the function tested gives
	for _, pair := range updateCellTests {

		v := UpdateCell(pair.input.b, pair.input.r, pair.input.c, pair.input.code, pair.input.numDays)
		if v != pair.value {
			t.Error("For", pair.input,
				"Expected", pair.value,
				"got", v)
		}
	}

}

// Testing  UpdateIntoNewBoard function
//Create type testpair
type UpdateIntoNewBoardInput struct {
	b       Board
	code    int
	numDays int
}

type updateintonewboardpair struct {
	input UpdateIntoNewBoardInput
	value Board
}

var updateIntoNewBoardTests = []updateintonewboardpair{
	{UpdateIntoNewBoardInput{Board{{{"S", 0}, {"I", 0}, {"S", 0}}, {{"S", 0}, {"S", 0}, {"S", 0}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}, 11, 1}, Board{{{"I", 1}, {"I", 0}, {"I", 1}},
		{{"I", 1}, {"I", 1}, {"I", 1}},
		{{"S", 0}, {"S", 0}, {"S", 0}}}}}

func Test_UpdateIntoNewBoard(t *testing.T) {

	for _, pair := range updateIntoNewBoardTests {

		v := UpdateIntoNewBoard(pair.input.b, pair.input.code, pair.input.numDays)
		if !reflect.DeepEqual(v, pair.value) {
			t.Error("For", pair.input,
				"Expected", pair.value,
				"got", v)
		}
	}

}
