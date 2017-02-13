package main

import "fmt"

func main() {
	a := make([]int, 3)
	a[0] = 1
	fmt.Println(a)
	b := append(a, 6)
	fmt.Println(b)
}
