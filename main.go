package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reading request
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	size := in.Text()
	fieldSize := strings.Split(size, " ")
	lines, err := strconv.Atoi(fieldSize[0])
	if err != nil {
		fmt.Println("Ошибка конвертации.", err)
	}
	points, err := strconv.Atoi(fieldSize[1])
	if err != nil {
		fmt.Println("Ошибка конвертации.", err)
	}
	allLines := make(map[int][]rune)
	for i := 0; i < lines; i++ {
		in.Scan()
		l := in.Text()
		allLines[i] = []rune(l)
	}
	// relocation slashes
	for l := 0; l < lines; l++ {
		p := allLines[l]
		for i := points - 1; i >= 0; i-- {
			if string(p[i]) == "\\" || string(p[i]) == "/" {
				o := allLines[l-1][i]
				allLines[l-1][i] = p[i]
				p[i] = o
			}
		}
	}

	// reverses and output
	revMap(allLines)
	for i := 0; i < len(allLines); i++ {
		revArray(allLines[i])
		fmt.Println(string(allLines[i]))
	}
}

// reverse array
func revArray(r []rune) []rune {
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	return r
}

// reverse map
func revMap(m map[int][]rune) map[int][]rune {
	for first, last := 0, len(m)-1; first < last; first, last = first+1, last-1 {
		m[first], m[last] = m[last], m[first]
	}
	return m
}
