package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	x := strings.Split(argsWithoutProg[0], "")
	y := strings.Split(argsWithoutProg[1], "")

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

	fmt.Println(b)
	fmt.Println(c)
}
