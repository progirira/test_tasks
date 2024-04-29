package main

import (
	"fmt"
	"sort"
)

func ifEachElementIsDevided(numbers []int, divisor int) bool {
	for _, num := range numbers {
		if num%divisor != 0 {
			return false
		}
	}
	return true
}

func findCommonDivisors(numbers []int) []int {
	sort.Ints(numbers)
	divisors := []int{}
	for i := 2; i <= numbers[0]; i++ {
		if ifEachElementIsDevided(numbers, i) {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func main() {
	var numbers = []int{42, 12, 18}
	fmt.Println(numbers)
	fmt.Println(findCommonDivisors(numbers))
}
