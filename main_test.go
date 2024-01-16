package main

import (
	"fmt"
	"math/rand"
	"testing"
)

//you can run all you tests by typing
//go test -v
//in the terminal window

// this is a utility function for testing
// it will return a random square on the grid
// it does not keep track of any previously returned grids
func getRandomGridSquare() []int {

	row := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	column := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	return []int{rand.Intn(len(row)) + 1, rand.Intn(len(column)) + 1}

}

//these are the two tests we have for our functions in main
//the purpose of tests is to mimic interaction with our code
//there is no "user input" - the test is the calling code

// here is an example of a failing test - what do we need to do to fix it?
func TestCreateGrid(t *testing.T) {
	// Arragne
	// needs nothing

	// Act
	grid := CreateGrid()

	// Assert - custom assert
	assertGridIsCorrectSize(t, grid, 7, 7)
}

//one good place to start here is by using our utility function
//to target a random grid square rather than 1,1 co-ordinates every time

func TestPlayerOneTakingShot(t *testing.T) {

	// Arrange
	grid := CreateGrid()

	// Act - Code Under Test - Production Code
	shotResult := PlayerOneTurn(grid, []int{1, 1})

	// Assert - check the result is what we
	if shotResult != false {
		t.Error("Shot should be false!")
	}
}

func TestPlayerTwoTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerTwoTurn(grid, []int{1, 1})
	if shotResult != true {
		t.Error("Shot should be true!")
	}
}

//other tests here that fail

// sometimes we write tests to test our own functions.
func TestGetRandomGridSquare(t *testing.T) {
	gridSquare := getRandomGridSquare()

	//literally only exists here to show you the output
	//should not exist in a real test
	fmt.Println(gridSquare)

	//poor test making use of magic numbers
	//you should probably re-write it
	if gridSquare[0] <= 0 || gridSquare[0] >= 10 {
		t.Error("Grid square row should be >0 and <10, but got: ", gridSquare[0])
	}

	if gridSquare[1] <= 0 || gridSquare[1] >= 10 {
		t.Error("Grid square column should be >0 and <10, but got: ", gridSquare[1])
	}
}

func TestPlaceAShip(t *testing.T) {
	// Arange
	grid := CreateGrid()

	// Act
	desiredCol := 3
	desiredRow := 5
	updatedGrid := PlaceShip(grid, desiredCol, desiredRow)

	// by here we assukme the ship has been placed on the grid

	// Assert
	// a ship is placed
	actual := updatedGrid[3][5]
	want := "S"
	if actual != want {
		t.Error("Ship was not placed at col 3, row 5")
	}
}

func assertGridIsCorrectSize(t *testing.T, grid [7][7]string, expectedRows int, expectedCols int) {
	gridSizeCols := len(grid)
	if gridSizeCols != expectedCols {
		t.Error("Grid is wrong size. Expected max size of 7, got: ", gridSizeCols)
	}

	gridSizeRows := len(grid[0])
	if gridSizeRows != expectedRows {
		t.Error("Grid has wrong number of rows, wanted 7 but was", gridSizeRows)
	}
}

func TestCannotPlaceShipOutsideGrid(t *testing.T) {

	// Arrange
	grid := CreateGrid()

	// Act
	// Trying to place a ship outside of the grid
	updatedGrid := PlaceShip(grid, 8, 8)

	// Assert
	for a := 0; a < 7; a++ {
		for b := 0; b < 7; b++ {
			if updatedGrid[a][b] != "" {
				t.Error("Ship should not be placed outside of the grid!!")
			}
		}
	}

}

func TestCannotPlaceTenthShip(t *testing.T) {

	// Arrange

	grid := CreateGrid()

	// Act
	for i := 0; i < maxShips; i++ { // create a max ship const in main for this 00:53 -- outside of placeship func
		grid = PlaceShip(grid, i, i)
	}

	updatedGrid := PlaceShip(grid, 0, 0)

	// Assert

	for a := 0; a < 7; a++ {
		for b := 0; b < 7; b++ {
			if updatedGrid[a][b] != "" {
				t.Errorf("Cant place more then 9 ships!!!!")
			}
		}
	}
}
