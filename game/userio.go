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
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func displayState(g game) {
	fmt.Println("Here's how the board looks like :")
	for i := 0; i < 9; i += 3 {
		fmt.Println("\t" + g.Board[i] + "|" + g.Board[i+1] + "|" + g.Board[i+2])
		if i < 5 {
			fmt.Println("\t" + "-+-+-")
		}
	}
	fmt.Println("")
	fmt.Println("======================================")
}

func printDraw() {
	fmt.Println("Its a draw!")
}

func printWinner(winner string) {
	fmt.Println("The winner is: ", winner)
}

func askUserMove() {
	fmt.Println("Please select a move [0 - 8]:")
}

func getUserMove(g game) int {
	pos := -1
	for true {
		userMove := readInputLine()
		pos, err := strconv.Atoi(userMove)
		if err != nil {
			fmt.Println("Invalid input please try again")
			continue
		}
		if pos < 0 || pos > 8 {
			fmt.Println("Please select a valid input")
			continue
		}
		if g.Board[pos] != " " {
			fmt.Println("Please select an empty space")
			continue
		}
		return pos
	}
	return pos
}

func askFirstInput() string {
	fmt.Println("Please chose your side. Note that x goes first.")
	return readInputLine()
}

func readInputLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text[:1]
}
