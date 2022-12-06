package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var stacks = [][]string{
	{"W", "R", "F"},
	{"T", "H", "M", "C", "D", "V", "W", "P"},
	{"P", "M", "Z", "N", "L"},
	{"J", "C", "H", "R"},
	{"C", "P", "G", "H", "Q", "T", "B"},
	{"G", "C", "W", "L", "F", "Z"},
	{"W", "V", "L", "Q", "Z", "J", "G", "C"},
	{"P", "N", "R", "F", "W", "T", "V", "C"},
	{"J", "W", "H", "G", "R", "S", "V"}}

// Move from input text
type Move struct {
	amount,
	from,
	to int
}

func main() {
	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	fileString := string(file)
	idx := strings.Index(fileString, "m")
	fileString = fileString[idx:]

	fileString = strings.ReplaceAll(fileString, "move", "")
	fileString = strings.ReplaceAll(fileString, "from", "")
	fileString = strings.ReplaceAll(fileString, "to", "")

	stringArr := strings.Split(fileString, "\n")

	moves := make([]Move, 0)
	newArr := [][]string{}
	for _, strArr := range stringArr {
		s := strings.TrimSpace(strArr)
		a := strings.Split(s, " ")
		for idx, val := range a {
			if val == "" {
				a = append(a[:idx], a[idx+1:]...)
			}
		}
		// fmt.Printf("%#v", a)
		newArr = append(newArr, a)
	}
	// fmt.Printf("%#v", newArr)

	for _, arr := range newArr {
		amt, err := strconv.Atoi(arr[0])
		from, err := strconv.Atoi(arr[1])
		to, err := strconv.Atoi(arr[2])

		if err != nil {
			panic(err)
		}
		moves = append(moves, Move{amount: amt, from: from, to: to})
	}

	for _, move := range moves {
		fr := stacks[move.from-1]
		t := stacks[move.to-1]

		for i := len(fr) - 1; i >= 0 && move.amount > 0; i-- {
			val := fr[i]
			t = append(t, val)
			fr = fr[:i]
			move.amount--
		}
		stacks[move.from-1] = fr
		stacks[move.to-1] = t

		// fmt.Println(stacks[move.from-1], stacks[move.to-1], move.amount)
		// fmt.Printf("%+v \n", stacks)
	}
	str := strings.Builder{}
	for _, val := range stacks {
		str.WriteString(val[0])
	}
	fmt.Println(str.String())
	// FZMRCPTMH -- incorrect will review tomorrow
}
