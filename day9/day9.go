package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x,
	y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func sign(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
}

func (p point) follow(h point) point {
	diffx, diffy := h.x-p.x, h.y-p.y
	// check distance whether to follow
	if abs(diffx) <= 1 && abs(diffy) <= 1 {
		return p
	}
	return point{p.x + sign(diffx), p.y + sign(diffy)}
}

func main() {
	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	var h, t point
	visited := map[point]struct{}{}
	for _, move := range strings.Split((string(file)), "\n") {
		m := strings.Split(move, " ")

		dir := string(m[0])
		num, _ := strconv.Atoi(string(m[1]))
		fmt.Println(num)
		for i := 0; i < num; i++ {
			switch dir {
			case "L":
				h.x--
			case "R":
				h.x++
			case "U":
				h.y++
			case "D":
				h.y--
			}
			t = t.follow(h)
			visited[t] = struct{}{}
		}
	}
	// part 1
	fmt.Println(len(visited))
}
