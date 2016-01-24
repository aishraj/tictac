// Copyright © 2016 Aish Raj Dahal <name>@aishraj.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package game

import (
	"fmt"
)

type game struct {
	Board         []string
	CurrentPlayer string
}

//StartGame is the starting point of a new Tic Tac Toe game instance.
func StartGame() {
	currentGame := newGame()
	userSymbol := "x"
	for !isOver(currentGame) {
		displayState(currentGame)
		fmt.Println("Please select a move [0 - 8]:")
		userPosition := getUserMove(currentGame)
		currentGame.Board[userPosition] = userSymbol
		if currentGame.CurrentPlayer == "x" {
			currentGame.CurrentPlayer = "o"
		} else {
			currentGame.CurrentPlayer = "x"
		}
		if isOver(currentGame) {
			break
		}
		minScore := betaMax
		var nextState game
		possibleStates := getPossibleStates(currentGame)
		for _, g := range possibleStates {
			score := minimax(g)
			if score <= minScore {
				minScore = score
				nextState = g
			}
		}
		currentGame = nextState
	}
	displayState(currentGame)
	if winner := getWinner(currentGame); winner == "draw" {
		fmt.Println("Its a draw")
	} else {
		fmt.Println("The winner is", winner)
	}
}

func getPossibleStates(g game) []game {
	var games []game
	var newPlayer string

	if g.CurrentPlayer == "x" {
		newPlayer = "o"
	} else {
		newPlayer = "x"
	}

	for i := 0; i < 9; i++ {
		if g.Board[i] == " " {
			nextBoard := make([]string, len(g.Board))
			copy(nextBoard, g.Board)
			nextBoard[i] = g.CurrentPlayer
			successor := game{Board: nextBoard, CurrentPlayer: newPlayer}
			games = append(games, successor)
		}
	}
	return games
}

func isOver(g game) bool {
	return getWinner(g) != ""
}

func getWinner(g game) string {
	gameLines := getGameLines(g)
	for _, line := range gameLines {
		if currentLine := winningLine(line); currentLine != "" {
			return currentLine
		}
	}
	if isFull(g) {
		return "draw"
	}
	return ""
}

func winningLine(line []string) string {
	if line[0] == line[1] && line[1] == line[2] {
		if line[0] == "x" {
			return "x"
		} else if line[0] == "o" {
			return "o"
		}
	}
	return ""
}

func isFull(g game) bool {
	for _, v := range g.Board {
		if v == " " {
			return false
		}
	}
	return true
}

func getGameLines(g game) [][]string {
	var lines [][]string
	for i := 0; i < 3; i++ {
		row := make([]string, 3)
		row[0] = g.Board[i*3]
		row[1] = g.Board[i*3+1]
		row[2] = g.Board[i*3+2]
		lines = append(lines, row)
	}
	for i := 0; i < 3; i++ {
		column := make([]string, 3)
		column[0] = g.Board[i]
		column[1] = g.Board[i+3]
		column[2] = g.Board[i+6]
		lines = append(lines, column)
	}
	leftDiagonal := []string{g.Board[0], g.Board[4], g.Board[8]}
	rightDiagonal := []string{g.Board[2], g.Board[4], g.Board[6]}
	lines = append(lines, leftDiagonal)
	lines = append(lines, rightDiagonal)
	return lines
}

func newGame() game {
	gameBoard := make([]string, 9)
	for i := 0; i < 9; i++ {
		gameBoard[i] = " "
	}
	currentPlayer := "x"
	return game{Board: gameBoard, CurrentPlayer: currentPlayer}
}
