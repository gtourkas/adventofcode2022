package adventofcode2022

import "strings"

func calcTotalScorePart1(lines []string) int {

	symbolToShape := map[string]string{
		"A": "Rock",
		"X": "Rock",
		"B": "Paper",
		"Y": "Paper",
		"C": "Scissors",
		"Z": "Scissors",
	}

	shapeScore := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	totalScore := 0
	for _, l := range lines {
		parts := strings.Split(l, " ")
		opponentShape := symbolToShape[parts[0]]
		myShape := symbolToShape[parts[1]]

		resultScore := 0
		if myShape == opponentShape {
			resultScore = 3
		}
		if (myShape == "Rock" && opponentShape == "Scissors") ||
			(myShape == "Scissors" && opponentShape == "Paper") ||
			(myShape == "Paper" && opponentShape == "Rock") {
			resultScore = 6
		}

		totalScore += resultScore + shapeScore[myShape]
	}

	return totalScore
}

func calcTotalScorePart2(lines []string) int {

	symbolToShape := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	symbolToResult := map[string]string{
		"X": "Lose",
		"Y": "Draw",
		"Z": "Win",
	}

	shapeScore := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	totalScore := 0
	for _, l := range lines {
		parts := strings.Split(l, " ")
		opponentShape := symbolToShape[parts[0]]
		result := symbolToResult[parts[1]]

		var myShape string
		resultScore := 0
		if result == "Lose" {
			if opponentShape == "Rock" {
				myShape = "Scissors"
			}
			if opponentShape == "Scissors" {
				myShape = "Paper"
			}
			if opponentShape == "Paper" {
				myShape = "Rock"
			}
		}
		if result == "Draw" {
			resultScore = 3
			myShape = opponentShape
		}
		if result == "Win" {
			resultScore = 6
			if opponentShape == "Rock" {
				myShape = "Paper"
			}
			if opponentShape == "Scissors" {
				myShape = "Rock"
			}
			if opponentShape == "Paper" {
				myShape = "Scissors"
			}
		}

		totalScore += resultScore + shapeScore[myShape]
	}

	return totalScore
}
