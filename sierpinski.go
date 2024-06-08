package main

import (
	"fmt"
	"strings"
)

func sierpinski(n int) []string {
	if n == 0 {
		return []string{"*"}
	}
	previous := sierpinski(n - 1)
	size := len(previous)
	result := make([]string, size*2)
	space := strings.Repeat(" ", size)
	for i := 0; i < size; i++ {
		result[i] = space + previous[i] + space
		result[size+i] = previous[i] + " " + previous[i]
	}
	return result
}

func main() {
	n := 4 // Increase this value to generate a larger triangle
	triangle := sierpinski(n)
	for _, line := range triangle {
		fmt.Println(line)
	}
}
