package main

import (
	"fmt"
	"math"
)

func ifNumIsSimple(number int) bool {
	sqrtNumber := int(math.Round(math.Sqrt(float64(number))))
	for divisor := 2; divisor <= sqrtNumber; divisor++ {
		if number%divisor == 0 {
			return false
		}
	}
	return true
}

func findSimpleNumbers(startNumber int, endNumber int) []int {
	simpleNumbers := []int{}
	for num := startNumber; num <= endNumber; num++ {
		if ifNumIsSimple(num) {
			simpleNumbers = append(simpleNumbers, num)
		}
	}
	return simpleNumbers
}

func main() {
	var start, end int
	fmt.Scanln(&start)
	fmt.Scanln(&end)
	fmt.Println(findSimpleNumbers(start, end))
}
