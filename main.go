package main

import "errors"

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

func PlayerOneTurn(playerTwoGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return false //shot missed
}

func PlayerTwoTurn(playerOneGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return true //shot hit
}

func CreateGrid() (grid [7][7]string) {
	// Initialise grid with an S at 2, 3
	for i := range grid {
		for j := range grid[i] {
			if i == 2 && j == 3 {
				grid[i][j] = "S"
			} else {
				grid[i][j] = ""
			}
		}
	}
	return grid
}

func countShips(grid [7][7]string) (int, error) {
	shipCount := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "S" {
				shipCount++
				if shipCount > 9 {
					return 0, errors.New("Cant place more than 9 ships.")
				}
			}
		}
	}
	return shipCount, nil
}

func placeShip(grid [7][7]string, coordinates []int) error {
	row, col := coordinates[0], coordinates[1]

	// Check coorrds within the grid
	if row < 0 || row >= 7 || col < 0 || col >= 7 {
		return errors.New("Cannot place ship outside the grid.")
	}

	// check if ship in coords
	if grid[row][col] == "S" {
		return errors.New("Cannot place ship on an occupied cell.")
	}

	// place ship at coords
	grid[row][col] = "S"
	return nil
}
