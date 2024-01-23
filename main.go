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

var errCannotPlaceMoreThanMaxShips = errors.New("you cant place more than maxShips")
var errCannotPlaceShipOutsideGrid = errors.New("cant place a ship outside of the grid")
var errCannotPlaceShipOnTopOfAnother = errors.New("cannot place a ship on top of another")

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
