package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	filestr := string(file)

	strarr := strings.Split(filestr, "\n")
	total := 0

	for _, str := range strarr {

		h2 := str[(len(str)+1)/2:]
		h1 := str[:(len(str))/2]

		m := map[rune]bool{}
		for _, char := range h1 {
			m[char] = true
		}

		for _, char := range h2 {
			if _, ok := m[char]; ok {
				total += priority(char)
				break
			}
		}
	}

	// fmt.Println(total)
	// 7821
	groupTotal := 0
	for i := 0; i < len(strarr); i += 3 {
		str1 := strarr[i]
		str2 := strarr[i+1]
		str3 := strarr[i+2]

		m1 := map[rune]bool{}
		m2 := map[rune]bool{}

		for _, char := range str1 {
			m1[char] = true
		}

		for _, char := range str2 {
			if _, ok := m1[char]; ok {
				m2[char] = true
			}
		}
		// fmt.Println(m2)
		for _, char := range str3 {
			if _, ok := m2[char]; ok {
				groupTotal += priority(char)
				break
			}
		}
	}
	// fmt.Println(groupTotal)
	// 2752
}

func priority(char rune) int {
	if char-'a' < 0 {
		return int(52 + (char - 'a') + 7)
	}

	return int(char - 'a' + 1)

}
