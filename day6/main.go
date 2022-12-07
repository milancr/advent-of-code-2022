package main

import (
	"fmt"
	"io/ioutil"
)

// find start of packet marker --4 different chars
// report index of when the index of the last of the chars

// windowLen for start of packets = 4
// windowLen for start of message = 14
var windowLen int = 14

func main() {

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	inputstr := string(file)
	charArr := []string{}
	marker := 0
	for idx, char := range inputstr {
		if len(charArr) == windowLen {
			if checkDuplicates(charArr) {
				marker = idx
				fmt.Println(charArr)
				break
			}
			charArr = charArr[1:]
			charArr = append(charArr, string(char))

		}
		if len(charArr) < windowLen {
			charArr = append(charArr, string(char))
		}

	}

	fmt.Println(marker)
	// 1625
	// 2250
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
