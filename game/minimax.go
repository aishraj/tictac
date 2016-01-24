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

const alphaMin = -100
const betaMax = 100

func minimax(g game) int {
	alpha := alphaMin
	beta := betaMax
	if g.CurrentPlayer == "x" {
		return maximize(g, alpha, beta)
	}
	return minimize(g, alpha, beta)
}

func maximize(g game, alpha int, beta int) int {
	if isOver(g) {
		return calculateScore(g)
	}
	v := alphaMin
	possibleStates := getPossibleStates(g)
	for _, nextState := range possibleStates {
		roundScore := minimize(nextState, alpha, beta)
		if roundScore >= v {
			v = roundScore
		}
		if v >= beta {
			return v
		}
		if v > alpha {
			alpha = v
		}
	}
	return v
}

func minimize(g game, alpha int, beta int) int {
	if isOver(g) {
		return calculateScore(g)
	}
	v := betaMax
	possibleStates := getPossibleStates(g)
	for _, nextState := range possibleStates {
		roundScore := maximize(nextState, alpha, beta)
		if roundScore <= v {
			v = roundScore
		}
		if v <= alpha {
			return v
		}
		if v < beta {
			beta = v
		}
	}
	return v
}

func calculateScore(g game) int {
	winner := getWinner(g)
	if winner == "draw" {
		return 0
	} else if winner == "x" {
		return 1
	} else if winner == "o" {
		return -1
	}
	return 0
}
