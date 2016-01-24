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
	"testing"
)

func TestPosition(t *testing.T) {
	board := []string{" ", " ", " ", " ", "x", " ", " ", " ", " "}
	expected := []int{0, 1, 0, 1, 1, 0, 1, 0}
	var actual []int
	for i := 0; i < 9; i++ {
		if board[i] == " " {
			board[i] = "o"
			g := game{Board: board, CurrentPlayer: "x"}
			score := minimax(g)
			actual = append(actual, score)
			board[i] = " "
		}
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error("Expected anr actual score not equal")
		}
	}
}
