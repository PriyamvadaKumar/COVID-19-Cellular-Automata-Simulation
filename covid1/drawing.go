package main

import (
	"canvas"
	"fmt"
	"image"
)

// DrawBoards is to generate gif images from a slice of Boards
func DrawBoards(boards []*Board, canvasWidth, cellWidth, q int) []image.Image {

	images := make([]image.Image, len(boards))
	for i, board := range boards {
		pic := DrawBoard(board, canvasWidth, cellWidth, q)
		images[i] = canvas.GetImage(&pic)
	}
	return images
}

//DrawBoard is to generate pic from Board
func DrawBoard(board *Board, canvasWidth, cellWidth, q int) canvas.Canvas {
	height := len(board.cells) * cellWidth
	width := len(board.cells) * cellWidth
	pic := canvas.CreateNewCanvas(width, height)
	// Color definitions

	/*
		blue := canvas.MakeColor(0, 0, 255)
		red := canvas.MakeColor(255, 0, 0)
		green := canvas.MakeColor(0, 255, 0)
		black := canvas.MakeColor(0, 0, 0)
	*/
	// Set up canvas

	pic.SetFillColor(canvas.MakeColor(0, 0, 0))
	pic.ClearRect(0, 0, canvasWidth, canvasWidth)
	pic.Fill()
	for i, cell := range board.cells {
		// Create Color based on proportion of classes
		S, I, R := ClassSum(cell, q)
		n1, n2, n3 := uint8(int(S*float64(255))), uint8(int(I*float64(255))), uint8(int(R*float64(255)))
		fmt.Println(n1, n2, n3)
		cc := canvas.MakeColor(n1, n2, n3)
		pic.SetFillColor(cc)
		//Coordinates of corner from which to draw the rectangle
		x := i * cellWidth
		y := 2 * cellWidth
		pic.ClearRect(x, y, x+cellWidth, y+cellWidth)
		pic.Fill()

	}

	return pic
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
