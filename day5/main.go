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

// Move is the instruction per line of  input text
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

		newArr = append(newArr, a)
	}

	for _, arr := range newArr {
		amt, err := strconv.Atoi(arr[0])
		from, err := strconv.Atoi(arr[1])
		to, err := strconv.Atoi(arr[2])

		if err != nil {
			panic(err)
		}
		moves = append(moves, Move{amount: amt, from: from, to: to})
	}

	// ! bug starting here
	// tried making copies using copy method and appending
	stack1 := make([][]string, 0)
	stack2 := make([][]string, 0)
	for _, arr := range stacks {
		stack1 = append(stack1, arr)
	}
	for _, arr := range stacks {
		stack2 = append(stack2, arr)
	}

	// calling crane 1 before crane2 causes crane2 to return the wrong answer
	// since I'm making copies of global variable stacks the operations should be executed on the copies
	// and not affect the original
	// not sure why one would affect the other, any suggestions appreciated
	fmt.Println(crane1(moves, stack1))
	// correct: CVCWCRTVQ
	fmt.Println(crane2(moves, stack2))
	// correct: CNSCZWLVT

}

func crane1(moves []Move, s1 [][]string) string {
	stack1 := make([][]string, 0)
	for _, arr := range stacks {
		stack1 = append(stack1, arr)
	}
	for _, move := range moves {
		from := stack1[move.from-1]
		to := stack1[move.to-1]

		for i := len(from) - 1; i >= 0 && move.amount > 0; i-- {
			val := from[i]
			to = append(to, val)
			from = from[:i]
			move.amount--
		}
		stack1[move.from-1] = from
		stack1[move.to-1] = to

	}
	str := strings.Builder{}
	for _, val := range stack1 {
		str.WriteString(val[len(val)-1])
	}
	return str.String()
}

func crane2(moves []Move, s2 [][]string) string {
	for _, move := range moves {
		from := s2[move.from-1]
		to := s2[move.to-1]

		i := len(from) - move.amount
		if i < 0 {
			i = 0
		}
		crates := from[i:]
		from = from[:i]
		to = append(to, crates...)

		s2[move.from-1] = from
		s2[move.to-1] = to

	}

	str := strings.Builder{}
	for _, val := range s2 {
		str.WriteString(val[len(val)-1])
	}
	return str.String()
}
