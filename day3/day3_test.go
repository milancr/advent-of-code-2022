package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

var strarr []string

func Benchmark_Serial(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
	}
}

func Benchmark_Parallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		ch := make(chan int)
		for i := 0; i < len(strarr); i += 3 {
			go findBadgeCode(strarr[i], strarr[i+1], strarr[i+2], ch)
		}

		for i := 0; i < len(strarr); i += 3 {
			sum += <-ch
		}

	}

}

func init() {
	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	strarr = strings.Split(string(file), "\n")

}
