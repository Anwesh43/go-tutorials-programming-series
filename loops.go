package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		if (i+1)%2 == 0 {
			continue
		}
		fmt.Println(i + 1)
		sum = sum + (i + 1)
		if i+1 == 5 {
			break
		}
	}
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		if (i+1)%2 == 1 {
			continue
		}
		a[i] = (i + 1) * 2
		if i+1 == 5 {
			break
		}
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
