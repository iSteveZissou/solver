package Solver

import (
	"log"
	"time"
)

type Cell struct {
	row      int
	col      int
	value    int
	solution int
	colgroup bool
}

type Grid struct {
	cells [9][9]Cell
}

type NewDataStruct struct {
	Test int
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (grid *Grid) isValid(c Cell, v int) bool {

	col := c.col
	row := c.row

	//check row
	for i := 0; i < 9; i++ {
		if grid.cells[row][i].value == v {
			// fmt.Println("found in row")
			return false
		}
	}

	//check column
	for i := 0; i < 9; i++ {
		if grid.cells[i][col].value == v {

			// fmt.Println("already here!")
			return false
		}

	}
	var nrow = c.row
	var ncol = c.col

	//check grid // redo this!!!!
	var x1 = 3 * (nrow / 3)
	var y1 = 3 * (ncol / 3)
	var x2 = x1 + 2
	var y2 = y1 + 2

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if grid.cells[i][j].value == v {
				// fmt.Println("found in grid")
				return false
			}
		}
	}
	return true
}

func (grid *Grid) solve(c Cell) bool {
	// defer timeTrack(time.Now(), "Solving")

	if c.value == 10 {
		return true

	}

	if c.value != 0 {
		// fmt.Println("the value before goin in", c.value)
		return grid.solve(grid.getNextCell(c))

	}
	// fmt.Println("here goes the first empty", c.row, c.col)

	// var x, y = grid.getNextCell(c)

	// if c.value > 0 {
	// 	grid.solve()

	// }

	for i := 1; i <= 9; i++ {

		var valid = grid.isValid(c, i)

		if !valid {
			// fmt.Println("NOT VALID")
			continue
		}

		grid.cells[c.row][c.col].value = i
		// fmt.Println("I am cell ", c.row, c.col)
		// fmt.Println("Current value = ", grid.cells[c.row][c.col].value)

		var solved = grid.solve(grid.getNextCell(c))

		if solved {
			return true
		}
		grid.cells[c.row][c.col].value = 0

	}
	// fmt.Println("HHHHHHHHHJDDDDDDDDDDDDDDDDDDDDDDD")
	return false
}

func (grid *Grid) getNextCell(c Cell) Cell {

	var test Cell
	test.value = 10

	var row = c.row
	var col = c.col

	// fmt.Println("This goes in tooo", row, col)
	col++

	if col > 8 {
		col = 0
		row++
	}

	if row > 8 {
		// fmt.Println("This is the end")
		return test
	}

	// var curRow = c.row
	// var curCol = c.col
	// fmt.Println("return this cell:", row, col)
	return grid.cells[row][col]

}

//"newsolver"
func NewSolver(puzzle [9][9]int) [9][9]int {

	grid := new(Grid)
	var b [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid.cells[i][j].value = puzzle[i][j]
			grid.cells[i][j].row = i
			grid.cells[i][j].col = j

		}
	}

	if grid.solve(grid.cells[0][0]) {

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				b[i][j] = grid.cells[i][j].value

			}
		}
		return b
	}

	return b

}
