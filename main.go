package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	size := in.Text()
	fieldSize := strings.Split(size, " ")
	//fmt.Println(fieldSize)
	lines, err := strconv.Atoi(fieldSize[0])
	if err != nil {
		fmt.Println("Ошибка конвертации.", err)
	}
	allLines := make(map[int]string)
	for i := 0; i < lines; i++ {
		in.Scan()
		l := in.Text()
		allLines[i] = l
	}
	//fmt.Println(allLines)
	points, err := strconv.Atoi(fieldSize[1])
	if err != nil {
		fmt.Println("Ошибка конвертации.", err)
	}
	for l := lines - 1; l >= 0; l-- {
		p := allLines[l]
		r := []rune(p)
		var res []rune
		for i := points - 1; i >= 0; i-- {
			res = append(res, r[i])
		}
		fmt.Println(string(res))
	}
}
