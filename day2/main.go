package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	// Lost ...
	Lost = 0
	// Tie ...
	Tie = 3
	// Win ...
	Win = 6
)

var score = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

// var score2 = map[string]int{"A": 1, "B": 2, "C": 3, "X": Win, "Y": Tie, "Z": Lost}
func main() {
	file, err := ioutil.ReadFile("./input.txt")

	inputStr := string(file)

	if err != nil {
		panic(err)
	}

	strArr := strings.Split(inputStr, "\n")
	total := 0
	for _, pair := range strArr {
		oppHand := pair[0]
		userHand := pair[1]
		score := getOutcome(string(oppHand), string(userHand))
		total += score
	}

	fmt.Println(total)
	// 14827

	totalp2 := 0
	for _, pair := range strArr {
		oppHand := pair[0]
		result := pair[1]
		score := getOutcome2(string(oppHand), string(result))
		totalp2 += score
	}

	fmt.Println(totalp2)
	// 13889
}

func getOutcome(input1, input2 string) int {
	if score[input1] == score[input2] {
		return score[input2] + Tie
	} else if (score[input1] == 1 && score[input2] == 2) || (score[input1] == 2 && score[input2] == 3) || (score[input1] == 3 && score[input2] == 1) {
		return score[input2] + Win
	} else {
		return score[input2] + Lost
	}
}

func getOutcome2(input1, input2 string) int {
	if input2 == "Z" {
		return Win + onWin(input1)
	}
	if input2 == "Y" {
		return Tie + score[input1]
	} else {
		return Lost + onLose(input1)
	}
}

func onWin(oppHand string) int {
	if oppHand == "A" {
		return score["Y"]
	} else if oppHand == "B" {
		return score["Z"]
	} else {
		return score["X"]
	}
}

func onLose(oppHand string) int {
	if oppHand == "A" {
		return score["Z"]
	} else if oppHand == "B" {
		return score["X"]
	} else {
		return score["Y"]
	}
}
