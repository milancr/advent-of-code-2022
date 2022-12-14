package main

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
