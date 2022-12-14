package main

import (
	"fmt"
	"io/ioutil"
)

// windowLen for start of packets = 4
// windowLen for start of message = 14
var windowLen = 4
var windowLen2 = 14

func main() {

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	p1 := processChars(windowLen, file)
	p2 := processChars(windowLen2, file)

	// part 1
	fmt.Println(p1)
	// part 2
	fmt.Println(p2)

}

func checkDuplicates(arr []string) bool {
	m := map[string]bool{}
	for _, char := range arr {
		if _, ok := m[char]; ok {
			return false
		}
		m[char] = true
	}

	return true
}

func processChars(length int, file []byte) int {
	charArr := []string{}
	marker := 0
	for idx, char := range string(file) {
		if len(charArr) == length {
			if checkDuplicates(charArr) {
				marker = idx

				break
			}
			charArr = charArr[1:]
			charArr = append(charArr, string(char))

		}
		if len(charArr) < length {
			charArr = append(charArr, string(char))
		}

	}
	return marker
}
