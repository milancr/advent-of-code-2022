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

	rope := [10]point{}
	visited := map[point]struct{}{{0, 0}: {}}

	for _, move := range strings.Split(string(file), "\n") {
		m := strings.Split(move, " ")
		dir := m[0]
		num, _ := strconv.Atoi(m[1])
		for i := 0; i < num; i++ {
			switch dir {
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			case "U":
				rope[0].y++
			case "D":
				rope[0].y--
			}
			for j := 1; j < 10; j++ {
				rope[j] = rope[j].follow(rope[j-1])
			}
			visited[rope[9]] = struct{}{}
		}
	}
	fmt.Println(len(visited))
}
