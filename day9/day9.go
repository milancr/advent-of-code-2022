package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	file, err := ioutil.ReadFile("./input.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	var h, t point
// 	visited := map[point]struct{}{}
// 	for _, move := range strings.Split((string(file)), "\n") {
// 		m := strings.Split(move, " ")

// 		dir := string(m[0])
// 		num, _ := strconv.Atoi(string(m[1]))
// 		fmt.Println(num)
// 		for i := 0; i < num; i++ {
// 			switch dir {
// 			case "L":
// 				h.x--
// 			case "R":
// 				h.x++
// 			case "U":
// 				h.y++
// 			case "D":
// 				h.y--
// 			}
// 			t = t.follow(h)
// 			visited[t] = struct{}{}
// 		}
// 	}
// // part 1
// 	fmt.Println(len(visited))
// }
