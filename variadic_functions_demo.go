package main

import "fmt"

func sumOfNNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	fmt.Println(sumOfNNumbers(10, 20, 30))
	fmt.Println(sumOfNNumbers(1, 2))
	fmt.Println(sumOfNNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumOfNNumbers())
}
