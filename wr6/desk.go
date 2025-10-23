package main

import (
	"fmt"
	"os"
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

func main() {
	var row, col int
	fmt.Print("Задай количство строк: ")
	fmt.Fscan(os.Stdin, &row)
	fmt.Print("Задай количство колонок: ")
	fmt.Fscan(os.Stdin, &col)
	fmt.Println(chessboard(row, col))
}
