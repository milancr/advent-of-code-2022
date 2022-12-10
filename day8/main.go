package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var (
	size = 0
	m    = make(map[point]uint8)
)

var (
	left  = point{0, -1}
	right = point{0, 1}
	up    = point{-1, 0}
	down  = point{1, 0}
)

func (p point) add(change point) (point, bool) {
	next := point{p.x + change.x, p.y + change.y}
	return next, next.x >= 0 && next.x < size && next.y >= 0 && next.y < size
}

func visibleFromEdge() map[point]bool {
	v := map[point]bool{}
	for i := 0; i < size; i++ {
		walk(v, point{i, 0}, right)
		walk(v, point{i, size - 1}, left)
		walk(v, point{0, i}, down)
		walk(v, point{size - 1, i}, up)
	}
	return v
}

func walk(visible map[point]bool, start, direction point) {
	visible[start] = true
	tallest := m[start]

	for curr, valid := start, true; valid; curr, valid = curr.add(direction) {
		if m[curr] > tallest {
			visible[curr] = true
			tallest = m[curr]
		}
		if m[curr] == 9 {
			break
		}
	}
}

func init() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var i int
	for i = 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j := 0; j < len(line); j++ {
			m[point{i, j}] = line[j] - '0'

		}
	}
	size = i

}

func view(start, direction point) int {
	score := 0
	for next, valid := start.add(direction); valid; next, valid = next.add(direction) {
		score++
		if m[next] >= m[start] {
			break
		}
	}
	return score
}

func scenicScore() map[point]int {
	s := map[point]int{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			p := point{i, j}
			s[p] = view(p, right) * view(p, left) * view(p, up) * view(p, down)
		}
	}
	return s
}

func main() {

	v := visibleFromEdge()
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if v[point{i, j}] {
				count++
			}
		}
	}
	// part 1
	fmt.Println(count)

	s := scenicScore()
	max := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			p := point{i, j}
			if s[p] > max {
				max = s[p]
			}
		}
	}
	// part2
	fmt.Println(max)

}
