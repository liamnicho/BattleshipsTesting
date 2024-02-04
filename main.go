package main

import (
	"errors"
)

/*
This game of battleships is very simple to start:
There are 2 players
Each player has a grid which is 7*7
Each player has 9 Battleships, each of which can occupy only one square on their grid
Each player can place their battleships anywhere on this grid
Players take it in turns to pick any grid square reference
If the player hits a battleship, then it is sunk, and the turn passes to the opponent
If the player misses a battleship then it is called a miss, and the turn passes to the opponent
The player to first sink all their opponent's battleships is the winner
*/

//All code in here is example code, you do not have to keep any of it.

const maxShips = 9

var shipsPlaced = 0
var (
	errCannotPlaceMoreThanMaxShips   = errors.New("you cant place more than maxShips")
	errCannotPlaceShipOutsideGrid    = errors.New("cant place a ship outside of the grid")
	errCannotPlaceShipOnTopOfAnother = errors.New("cannot place a ship on top of another")
)

func PlayerOneTurn(playerTwoGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return false //shot missed
}

func PlayerTwoTurn(playerOneGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return true //shot hit
}

func CreateGrid() (grid [7][7]string) {
	//this is a fixed array not a slice
	return [7][7]string{}
}

func PlaceShip(grid [7][7]string, col int, row int) ([7][7]string, error) {
	if shipsPlaced >= maxShips {
		return grid, errCannotPlaceMoreThanMaxShips
	}
	if col < 0 || col >= 7 || row < 0 || row >= 7 {
		return grid, errCannotPlaceShipOutsideGrid
	}
	//  checks if a ship is already on a grid so you cant place another
	shipPlaced := checkIfShipPlaced(grid, col, row)
	if shipPlaced {
		return grid, errCannotPlaceShipOnTopOfAnother
	}
	grid[col][row] = "S"
	shipsPlaced++
	return grid, nil
}

func checkIfShipPlaced(grid [7][7]string, col int, row int) bool {
	return grid[col][row] == "S"
}

// takeTurn - takes grid, row and col as input and places the ship opponent ship
// grid is provided by one player, and row/col will by provided by the opponent
// If the ship is hit it's marked "sunk"
// else if the ship is missed, it's marked "missed"

func takeTurn(grid [7][7]string, col int, row int) ([7][7]string, bool) {
	shotStatus := false
	if grid[row][col] == "S" { // Hit case
		grid[row][col] = "sunk"
		shotStatus = true
	} else if grid[row][col] == "" {
		grid[row][col] = "missed"
	}
	return grid, shotStatus
}

// countShipsInGrid returns the number of ships in the grid
func countShipsInGrid(grid [7][7]string) int {
	shipsCount := 0
	for row := 0; row < 7; row++ {
		for col := 0; col < 7; col++ {
			if grid[row][col] == "S" {
				shipsCount++
			}
		}
	}
	return shipsCount
}
