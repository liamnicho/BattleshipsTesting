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

	grid := CreateGrid()

	gridSizeCols := len(grid)
	if gridSizeCols != 7 {
		t.Error("expected size of 7, got: ", gridSizeCols)
	}
	gridSizeRows := len(grid[0])
	if gridSizeRows != 7 {
		t.Error("expected 7, got:", gridSizeRows)
	}
}

//one good place to start here is by using our utility function
//to target a random grid square rather than 1,1 co-ordinates every time

func TestPlayerOneTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerOneTurn(grid, []int{1, 1})
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

func TestPlayerOneTurn(t *testing.T) {
	// Create a grid
	playerTwoGrid := CreateGrid()
	playerTwoGrid[2][3] = "S"

	// shot that misses
	shotResult := PlayerOneTurn(playerTwoGrid, []int{1, 1})
	if shotResult != false {
		t.Error("Expected shot to miss, but it hit.")
	}

	// shot that hits
	shotResult = PlayerOneTurn(playerTwoGrid, []int{2, 3})
	if shotResult != false {
		t.Error("Expected shot to miss, but it hit.")
	}
}

func TestPlayerTwoTurn(t *testing.T) {
	// Create a grid for player one with a battleship
	playerOneGrid := CreateGrid()
	playerOneGrid[4][5] = "S"

	// shot that misses
	shotResult := PlayerTwoTurn(playerOneGrid, []int{1, 1})
	if shotResult != true {
		t.Error("Expected to hit, but it missed")
	}

	// Test a shot that hits
	shotResult = PlayerTwoTurn(playerOneGrid, []int{4, 5})
	if shotResult != true {
		t.Error("Expected to hit, but it missed")
	}
}

func TestCountShips(t *testing.T) {
	// create grid with battleships
	grid := CreateGrid()
	grid[2][3] = "S"
	grid[4][5] = "S"
	grid[1][1] = "S"

	// Test placing more than 9 ships
	for i := 0; i < 7; i++ {
		grid[0][i] = "S"
	}

	_, err := countShips(grid)
	if err == nil {
		t.Error("Expected an error for placing more than 9 ships.")
	}
}

func TestCanPlaceShip(t *testing.T) {
	// Create grid with ship at 2, 3
	grid := CreateGrid()

	// Check ship is placed
	if grid[2][3] != "S" {
		t.Errorf("Expected ship at (2, 3), got: %v", grid)
	}
}

func TestCannotPlaceShipOutsideGrid(t *testing.T) {
	// empty grid
	grid := CreateGrid()

	// place ship outside grid at coords 7, 8
	err := placeShip(grid, []int{7, 8})

	// check if error is returned
	if err == nil {
		t.Error("Expected an error for placing a ship outside the grid, but no error received.")
	}
}

func TestCannotPlaceTenthShip(t *testing.T) {
	// Create grid with nine ships placed
	grid := CreateGrid()
	for i := 0; i < 7; i++ {
		grid[i][0] = "S"
	}

	// placing tenth ship
	err := placeShip(grid, []int{0, 0})

	// check error is returned
	if err == nil {
		t.Error("Expected an error for placing the tenth ship, but no error received.")
	}
}
