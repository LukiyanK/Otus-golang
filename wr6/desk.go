package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func chessboard(sizeX int, sizeY int) string {
	var builder strings.Builder
	for row := 0; row < sizeX; row++ {
		for col := 0; col < sizeY; col++ {
			if (row+col)%2 == 0 {
				builder.WriteString(" ")
			} else {
				builder.WriteString("#")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func checkinput(userinput string) int {
	var tempnum int
	tempnum, err := strconv.Atoi(userinput)
	if err != nil {
		fmt.Println("Ошибка ввода: будет использовано значение по умолчанию")
		tempnum = 8
	}
	return tempnum
}

func main() {
	var input string
	var row, col int
	fmt.Print("Задай количество строк(по умолчанию 8): ")
	fmt.Fscan(os.Stdin, &input)
	row = checkinput(input)
	fmt.Print("Задай количество колонок(по умолчанию 8): ")
	fmt.Fscan(os.Stdin, &input)
	col = checkinput(input)
	fmt.Println(chessboard(row, col))
}
