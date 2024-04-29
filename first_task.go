package main

import "fmt"

func howManyComputers(number int) string {
	if number >= 5 && number <= 20 {
		return fmt.Sprintf("%d компьютеров", number)
	}
	remainder := number % 10
	switch remainder {
	case 1:
		return fmt.Sprintf("%d компьютер", number)
	case 2, 3, 4:
		return fmt.Sprintf("%d компьютера", number)
	default:
		return fmt.Sprintf("%d компьютеров", number)
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(howManyComputers(n))
}
