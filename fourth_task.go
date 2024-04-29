package main

import (
	"fmt"
)

func printTable(table [][]int) {
	for _, row := range table {
		for _, value := range row {
			if value == 0 {
				fmt.Printf(" \t")
				continue
			}
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

func printAMultiplicationTable(number int) {
	table := make([][]int, number+1)
	for i := range table {
		table[i] = make([]int, number+1)
	}
	for i := 1; i <= number; i++ {
		table[0][i] = i
		table[i][0] = i
	}

	for rowNum := 1; rowNum <= number; rowNum++ {
		for colNum := 1; colNum <= number; colNum++ {
			table[rowNum][colNum] = rowNum * colNum
		}
	}
	printTable(table)
}

func main() {
	var number int
	fmt.Scanln(&number)
	printAMultiplicationTable(number)
}
