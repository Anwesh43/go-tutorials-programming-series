package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
		sum = sum + (i + 1)
	}
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = (i + 1) * 2
	}
	fmt.Println(a)
	fmt.Println(sum)
	fmt.Println(10 * (10 + 1) / 2)
	sum1 := 0
	for _, val := range a {
		sum1 = sum1 + val
	}
	fmt.Println("sum of the first 10 even numbers are", sum1)
}
