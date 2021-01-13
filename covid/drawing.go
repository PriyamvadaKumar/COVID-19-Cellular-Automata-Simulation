package main

import (
	"canvas"
	"image"
)

// DrawBoards is to generate gif images from a slice of Boards
func DrawBoards(boards []Board, cellWidth int) []image.Image {

	images := make([]image.Image, len(boards))
	for i, board := range boards {
		pic := DrawBoard(board, cellWidth)
		images[i] = canvas.GetImage(&pic)
	}
	return images
}

//DrawBoard is to generate pic fro a Game Board
func DrawBoard(board Board, cellWidth int) canvas.Canvas {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := canvas.CreateNewCanvas(width, height)
	// declare colors
	blue := canvas.MakeColor(0, 0, 255)
	yellow := canvas.MakeColor(255, 255, 0)
	red := canvas.MakeColor(255, 0, 0)
	green := canvas.MakeColor(0, 255, 0)
	black := canvas.MakeColor(0, 0, 0)
	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].status == "S" { //if status is susceptible S, then color with blue
				c.SetFillColor(blue)
			} else if board[i][j].status == "V" { //if status is vaccinated S, then color with yellow
				c.SetFillColor(yellow)
			} else if board[i][j].status == "I" { // if status is infected I,  then colour with red
				c.SetFillColor(red)

			} else if board[i][j].status == "R" { // if status is R  them colour with green
				c.SetFillColor(green)
			} else if board[i][j].status == "D" { // if status is D  them colour with black
				c.SetFillColor(black)

			} else {
				panic("Error: Out of range value " + board[i][j].status + " in board when drawing board.") // for panic error
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	return c
}

// DrawGridLines si to draw grid lines
func DrawGridLines(pic canvas.Canvas, cellWidth int) {
	w, h := pic.Width(), pic.Height()
	// first, draw vertical lines
	for i := 1; i < w/cellWidth; i++ {
		y := i * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	// next, draw horizontal lines
	for j := 1; j < h/cellWidth; j++ {
		x := j * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	pic.Stroke()
}
