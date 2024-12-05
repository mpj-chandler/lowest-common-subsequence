package main

import (
	"fmt"
	"os"
	"strings"
)

func construct_lowest_common_subsequence_table(x []string, y []string) [][]string {
	m := len(x)
	n := len(y)

	b := make([][]string, m)
	c := make([][]int8, m+1)

	for i := range b {
		b[i] = make([]string, n)
	}

	for i := range c {
		c[i] = make([]int8, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if x[i] == y[j] {
				c[i+1][j+1] = c[i][j] + 1
				b[i][j] = "UL"
			} else if c[i][j+1] >= c[i+1][j] {
				c[i+1][j+1] = c[i][j+1]
				b[i][j] = "U"
			} else {
				c[i+1][j+1] = c[i+1][j]
				b[i][j] = "L"
			}
		}
	}

	return b
}

func print_lowest_common_subsequence(b [][]string, x []string, i int, j int) {
	if i < 0 || j < 0 {
		return
	}
	if b[i][j] == "UL" {
		print_lowest_common_subsequence(b, x, i-1, j-1)
		fmt.Print(x[i])
	} else if b[i][j] == "U" {
		print_lowest_common_subsequence(b, x, i-1, j)
	} else {
		print_lowest_common_subsequence(b, x, i, j-1)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]

	x := strings.Split(argsWithoutProg[0], "")
	y := strings.Split(argsWithoutProg[1], "")

	b := construct_lowest_common_subsequence_table(x, y)

	print_lowest_common_subsequence(b, x, len(x)-1, len(y)-1)
}
