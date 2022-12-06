package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	strArr := strings.Split(string(file), "\n")

	arr := [][]string{}
	for _, str := range strArr {
		s := strings.Split(str, ",")

		arr = append(arr, s)
	}

	total := 0
	total2 := 0
	for _, pair := range arr {
		splitPair1 := strings.Split(pair[0], "-")
		elf1Start, err := strconv.Atoi(splitPair1[0])
		elf1End, err := strconv.Atoi(splitPair1[1])

		splitPair2 := strings.Split(pair[1], "-")
		elf2Start, err := strconv.Atoi(splitPair2[0])
		elf2End, err := strconv.Atoi(splitPair2[1])

		if err != nil {
			panic(err)
		}

		if covered(elf1Start, elf1End, elf2Start, elf2End) {
			total++
		}

		if covered2(elf1Start, elf1End, elf2Start, elf2End) {
			total2++
		}
	}

	fmt.Println(total)
	// 534
	fmt.Println(total2)
	//841
}

func covered(e1s, e1e, e2s, e2e int) bool {
	if e1s <= e2s && e2e <= e1e {
		return true
	}
	if e2s <= e1s && e2e >= e1e {
		return true
	}
	return false
}

func covered2(e1s, e1e, e2s, e2e int) bool {
	if e1e == e2e || e1s == e2s {
		return true
	}

	if (e1s <= e2s && e2s <= e1e) || (e2s <= e1s && e1s <= e2e) {
		return true
	}

	if (e1s <= e2e && e2e <= e1e) || (e2s <= e1e && e1e <= e2e) {
		return true
	}
	return false
}
