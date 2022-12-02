package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	str := string(file)

	if err != nil {
		panic(err)
	}

	arr := strings.Split(str, "\n")

	store := [][]string{}
	temp := []string{}

	for _, val := range arr {
		if val != "" {
			temp = append(temp, val)
		} else {
			store = append(store, temp)
			temp = []string{}
		}
	}

	if len(temp) != 0 {
		store = append(store, temp)
	}

	sums := []int{}
	for _, arr := range store {
		sum := 0
		for _, val := range arr {
			v, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			sum += v
		}
		sums = append(sums, sum)
		sum = 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	total := 0
	for _, v := range sums[:3] {
		total += v
	}
	// ans 1 | ans 2
	fmt.Println(sums[0], total)
}
